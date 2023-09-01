package udp

import (
	"net"
	"os"
	"sync"
)

var conn *net.UDPConn

func CreateConnection() *net.UDPConn {
	once.Do(func() {
		udpServer, err := net.ResolveUDPAddr("udp", ":1053")

		if err != nil {
			println("ResolveUDPAddr failed:", err.Error())
			os.Exit(1)
		}

		connection, err := net.DialUDP("udp", nil, udpServer)
		if err != nil {
			println("Listen failed:", err.Error())
			os.Exit(1)
		}
		conn = connection
		//send first message for handshake
		_, err = conn.Write([]byte("handshake"))
		if err != nil {
			println("Write data failed:", err.Error())
			os.Exit(1)
		}
	})
	return conn
}

func Receive(wg *sync.WaitGroup, handler func([]byte)) {
	for {
		received := make([]byte, 1024)
		_, err := conn.Read(received)
		if err != nil {
			println("Read data failed:", err.Error())
			break
		}
		handler(received)
	}
	if wg != nil {
		wg.Done()
	}
}

func Close() {
	defer conn.Close()
}
