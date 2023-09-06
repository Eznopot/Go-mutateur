package core

import (
	"encoding/json"
	"fmt"
	"go_mutateur/src/config"
	"go_mutateur/src/listener"
	"go_mutateur/src/prompter"
	"go_mutateur/src/udp"
	"net"

	hook "github.com/robotn/gohook"
)

var clientIndex int

// The function "handlerServer" receives UDP packets, prints the received data, and does not return any
// value.
//
// Args:
//   udpServer: udpServer is a variable of type net.PacketConn, which represents a connection for
// sending and receiving UDP packets. It is used to send and receive packets over a UDP network
// connection.
//   addr: The `addr` parameter is the network address of the remote client that sent the UDP packet.
// It contains information such as the IP address and port number of the client.
//   buf ([]byte): The `buf` parameter is a byte slice that represents the data received from the UDP
// server.
func handlerServer(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	fmt.Println("recus:", string(buf))
}

// The function `keyboardEventHandler` sends a keyboard event to a client over UDP.
//
// Args:
//   ev: The parameter "ev" is of type hook.Event. It represents the keyboard event that occurred.
func keyboardEventHandler(ev hook.Event) {
	toSend, err := json.Marshal(ev)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	udp.SendToClient(string(toSend), clientIndex-1)
}

// The function `keyboardEventHandlerForAll` sends a JSON-encoded keyboard event to all connected
// clients via UDP.
//
// Args:
//   ev: The parameter "ev" is of type hook.Event. It represents the keyboard event that occurred.
func keyboardEventHandlerForAll(ev hook.Event) {
	toSend, err := json.Marshal(ev)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	udp.SendToAllClient(string(toSend))
}

// The CoreServer function creates a UDP server, prompts the user to select a client to switch on, and
// handles events based on the selected client.
func CoreServer() {
	udp.CreateServer(config.GetConfig().Server.Port, handlerServer)
	for {
		client := udp.GetAllClientInfo()
		client = append(client, "All")
		client = append(client, "Exit")
		clientIndex = prompter.Select("Switch on client:", client)
		if clientIndex == len(client)-1 {
			break
		} else if clientIndex == len(client)-2 {
			listener.Event(keyboardEventHandler)
		} else {
			listener.Event(keyboardEventHandlerForAll)
		}
	}
	udp.Close()
}
