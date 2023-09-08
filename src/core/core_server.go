package core

import (
	"encoding/json"
	"fmt"
	"go_mutateur/src/config"
	"go_mutateur/src/listener"
	"go_mutateur/src/prompter"
	"go_mutateur/src/udp"
	"net"
	"os"
	"os/signal"

	hook "github.com/robotn/gohook"
)

var clientIndex int

func handlerServer(udpServer net.PacketConn, addr *net.Addr, packet udp.Packet) {
	//TODO: handle somthing
}

func keyboardEventHandler(ev hook.Event) {
	toSend, err := json.Marshal(ev)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	udp.SendToClient(string(toSend), "event", clientIndex-1)
}

func keyboardEventHandlerForAll(ev hook.Event) {
	toSend, err := json.Marshal(ev)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	udp.SendToAllClient(string(toSend), "event")
}

// The CoreServer function creates a UDP server, allows the user to select a client to switch on, and
// handles keyboard events for the selected client.
func CoreServer() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		udp.CloseServer()
		os.Exit(0)
	}()
	udp.CreateServer(config.GetConfig().Server.Port, handlerServer)
	for {
		client := udp.GetAllClientInfo()
		client = append(client, "All")
		client = append(client, "Exit")
		clientIndex = prompter.Select("Switch on client:", client)
		if clientIndex == -1 {
			continue
		} else if clientIndex == len(client) {
			break
		} else if clientIndex == len(client)-1 {
			listener.Event(keyboardEventHandlerForAll)
		} else {
			listener.Event(keyboardEventHandler)
		}
	}
	udp.CloseServer()
}
