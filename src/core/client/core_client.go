package core_client

import (
	"go_mutateur/src/config"
	"go_mutateur/src/listener"
	"go_mutateur/src/tui"
	"go_mutateur/src/udp"
	udp_client "go_mutateur/src/udp/client"
	"os"
	"os/signal"
)

func CoreClient() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		udp_client.CloseConnection()
	}()
	go udp_client.CreateConnection(config.GetConfig().Client.Address, config.GetConfig().Client.Port)
	tui.Init()
	tui.TuiCreateView("Select option", []string{"Exit"}, []rune{rune('q')}, process)
	go udp_client.Receive(nil, func(packet udp.Packet) {
		tui.AddLog(packet.Data)
		if packet.Type == "system" {
			switch data := packet.Data; data {
			case "close":
				tui.Close()
			default:
				break
			}
		} else {
			listener.Do(packet.Data)
		}
	})
	tui.Run()
}

func process(index int, main, secondary string, r rune) {
	if main == "Exit" {
		tui.Close()
		udp_client.CloseConnection()
	}
}
