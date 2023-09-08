package udp_client

import (
	"encoding/json"
	"fmt"
	"go_mutateur/src/udp"
	"log"
	"net"
	"sync"
)

var conn *net.UDPConn
var isConnClose = false
var once sync.Once

func CreateConnection(address, port string) *net.UDPConn {
	fmt.Println("Start new connection...")
	once.Do(func() {
		udpServer, err := net.ResolveUDPAddr("udp", address+":"+port)

		if err != nil {
			log.Fatal("ResolveUDPAddr failed:", err.Error())
		}

		connection, err := net.DialUDP("udp", nil, udpServer)
		if err != nil {
			log.Fatal("Listen failed:", err.Error())
		}
		conn = connection
		//send first message for handshake
		handshake := udp.Packet{
			Data: "handshake",
			Type: "system",
		}
		bytes, err := json.Marshal(handshake)
		if err != nil {
			log.Fatal("serialize failed:", err.Error())
		}
		_, err = conn.Write(bytes)
		if err != nil {
			log.Fatal("Write data failed:", err.Error())
		}
	})
	return conn
}

func clientPacketSystemHandler(data string) int {
	switch data := data; data {
	case "close":
		CloseConnection()
		return 1
	default:
		return 0
	}
}

func readFromConn() (udp.Packet, error) {
	var packet udp.Packet
	received := make([]byte, 2048)
	len, err := conn.Read(received)
	if isConnClose {
		return packet, nil
	} else if err != nil {
		println("Read data failed:", err.Error())
		return packet, err
	}
	received = received[:len]
	err = json.Unmarshal(received, &packet)
	if err != nil {
		log.Fatal("error on json:", err.Error())
		return packet, err
	}
	return packet, nil
}

func Receive(wg *sync.WaitGroup, handler func(udp.Packet)) {
	for !isConnClose {
		packet, err := readFromConn()
		if isConnClose {
			break
		} else if err != nil {
			continue
		}
		if packet.Type == "system" {
			if clientPacketSystemHandler(packet.Data) == 1 {
				return
			}
			continue
		}
		handler(packet)
	}
	if wg != nil {
		wg.Done()
	}
}

func CloseConnection() {
	packet := udp.Packet{
		Data: "close",
		Type: "system",
	}
	bytes, err := json.Marshal(packet)
	if err != nil {
		log.Fatal("error on json:", err.Error())
	}
	conn.Write(bytes)
	isConnClose = true
	defer conn.Close()
}
