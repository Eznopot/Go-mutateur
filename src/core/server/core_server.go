package core_server

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
	if packet.Type == "system" {
		switch data := packet.Data; data {
		case "handshake":
			index := len(udp_server.GetAllClientInfo()) - 1
			tui.InsertMenuOption((*addr).String(), index)
		case "close":
			index := tui.RemoveMenuOptionByString((*addr).String())
			if clientIndex == index {
				tui.FocusMenu()
			}
		}
	}
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

func onStart() {
	tui.FocusScreen()
	tui.AddLog("Open event stream with Client")
}

func onEnd() {
	tui.AddLog("Close event stream with Client")
	tui.FocusMenu()
}

func CoreServer() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		udp_server.CloseServer()
		tui.Close()
		os.Exit(0)
	}()
	udp_server.CreateServer(config.GetConfig().Server.Port, handlerServer)

	client := []string{"All", "Exit"}
	shortcuts := []rune{rune('a'), 'q'}
	tui.Init()
	tui.TuiCreateView("Choisissez une option:", client, shortcuts, process)
	tui.Run()
	udp_server.CloseServer()
}

func process(index int, title string, _ string, r rune) {
	if index == -1 {
		return
	} else if title == "Exit" {
		tui.Close()
	} else if title == "All" {
		go listener.Event(onStart, onEnd, keyboardEventHandlerForAll)
	} else {
		go listener.Event(onStart, onEnd, keyboardEventHandler)
	}
}
