package core_client

import (
	"bufio"
	"fmt"
	"go_mutateur/src/config"
	"go_mutateur/src/listener"
	"go_mutateur/src/tui"
	"os"
	"os/signal"

	"github.com/Eznopot/udp"
	udp_client "github.com/Eznopot/udp/client"
)

func warning() {
	if config.GetConfig().Client.Address == "127.0.0.1" || config.GetConfig().Client.Address == "localhost" {
		fmt.Println("Etes vous sur de vouloir vous connecter à un hote local ? Cela pourrais créer une boucle infini. (y/n)")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil || line != "y\n" {
			return
		}
	}
}

func CoreClient() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		udp_client.CloseConnection()
	}()

	warning()
	tui.Init()
	tui.TuiCreateView("Select option", []string{"Exit"}, []rune{rune('q')}, process)
	go tui.Run()
	udp_client.CreateConnection(config.GetConfig().Client.Address, config.GetConfig().Client.Port)
	udp_client.Receive(nil, func(packet udp.Packet) {
		tui.AddLog(packet.Data)
		if packet.Type == "system" {
			switch data := packet.Data; data {
			case "close":
				tui.Close()
			default:
				break
			}
		} else {
			if config.GetConfig().Developpement.Replication {
				listener.Do(packet.Data)
			}
		}
	})
}

func process(index int, main, secondary string, r rune) {
	if main == "Exit" {
		tui.Close()
		udp_client.CloseConnection()
	}
}
