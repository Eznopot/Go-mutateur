package core

import (
	"encoding/json"
	"fmt"
	"go_mutateur/src/listener"
	"go_mutateur/src/prompter"
	"go_mutateur/src/udp"
	"net"

	hook "github.com/robotn/gohook"
)

var clientIndex int

func handlerServer(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	fmt.Println("recus:", string(buf))
}

func keyboardEventHandler(ev hook.Event) {
	fmt.Println("keyboard event:", ev)
	toSend, err := json.Marshal(ev)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	udp.SendToClient(string(toSend), clientIndex-1)
}

func CoreServer() {
	udp.CreateServer(handlerServer)
	for {
		client := udp.GetAllClientInfo()
		clientIndex = prompter.Select("Switch on client:", client)
		if clientIndex == 0 {
			break
		}
		listener.Event(keyboardEventHandler)
	}
	//udp.Close()
}
