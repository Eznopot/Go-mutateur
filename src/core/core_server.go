package core

import (
	"encoding/json"
	"fmt"
	"go_mutateur/src/config"
	"go_mutateur/src/listener"
	"go_mutateur/src/tui"
	"go_mutateur/src/udp"
	udp_server "go_mutateur/src/udp/server"
	"net"
	"os"
	"os/signal"

	hook "github.com/robotn/gohook"
)

var clientIndex int

func handlerServer(udpServer net.PacketConn, addr *net.Addr, packet udp.Packet) {
	res, err := json.Marshal(packet)
	if err != nil {
		return
	}
	tui.AddLog(string(res))
}

func keyboardEventHandler(ev hook.Event) {
	toSend, err := json.Marshal(ev)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	tui.AddLog(string(toSend))
	udp_server.SendToClient(string(toSend), "event", clientIndex)
}

func keyboardEventHandlerForAll(ev hook.Event) {
	toSend, err := json.Marshal(ev)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	udp_server.SendToAllClient(string(toSend), "event")
}

// The CoreServer function creates a UDP server, allows the user to select a client to switch on, and
// handles keyboard events for the selected client.
// func CoreServer() {
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt)
// 	go func() {
// 		<-c
// 		udp_server.CloseServer()
// 		os.Exit(0)
// 	}()
// 	udp_server.CreateServer(config.GetConfig().Server.Port, handlerServer)
// 	for {
// 		client := udp_server.GetAllClientInfo()
// 		client = append(client, "All")
// 		client = append(client, "Exit")
// 		clientIndex = prompter.Select("Switch on client:", client)
// 		if clientIndex == -1 {
// 			continue
// 		} else if clientIndex == len(client) {
// 			break
// 		} else if clientIndex == len(client)-1 {
// 			listener.Event(keyboardEventHandlerForAll)
// 		} else {
// 			listener.Event(keyboardEventHandler)
// 		}
// 	}
// 	udp_server.CloseServer()
// }

func CoreServer() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		udp_server.CloseServer()
		tui.Close()
		os.Exit(0)
	}()
	udp_server.SetLogger(tui.AddLog)
	udp_server.CreateServer(config.GetConfig().Server.Port, handlerServer)

	client := udp_server.GetAllClientInfo()
	client = append(client, "All")
	client = append(client, "Exit")
	tui.Init()
	tui.TuiCreateView("Choisissez une option:", client, process)
	tui.Run()
	udp_server.CloseServer()
}

func process(index int, title string, _ string, r rune) {
	if index == -1 {
		return
	} else if title == "Exit" {
		tui.Close()
	} else if title == "All" {
		go listener.Event(keyboardEventHandlerForAll)
	} else {
		go listener.Event(keyboardEventHandler)
	}
}
