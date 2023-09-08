package udp

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"
)

var instance *net.PacketConn
var addrs []*net.Addr
var isServerClose = false

var once sync.Once

func serverPacketSystemHandler(data string, addr *net.Addr) int {
	switch data := data; data {
	case "handshake":
		fmt.Println("New client connected on:", *addr)
		if addrs == nil {
			var tmp []*net.Addr
			tmp = append(tmp, addr)
			addrs = tmp
			return 1
		}
		addrs = (append(addrs, addr))
		return 0
	case "close":
		fmt.Println("Client on:", *addr, "disconnected")
		for i, tmpAddr := range addrs {
			if (*addr).String() == (*tmpAddr).String() {
				addrs = append(addrs[:i], addrs[i+1:]...)
				break
			}
		}
		return 0
	default:
		return 0
	}
}
func listener(wg *sync.WaitGroup, handler func(net.PacketConn, *net.Addr, Packet)) {
	defer wg.Done()
	for !isServerClose {
		packet, addr, err := readFromSocket()
		if isServerClose {
			break
		} else if err != nil {
			println("error on socket", err.Error())
			continue
		}
		if packet.Type == "system" {
			if serverPacketSystemHandler(packet.Data, addr) == 1 {
				return
			}
			continue
		}
		handler(*instance, addr, packet)
	}
}

func readFromSocket() (Packet, *net.Addr, error) {
	var packet Packet
	buf := make([]byte, 2048)
	len, addr, err := (*instance).ReadFrom(buf)
	if isServerClose {
		return packet, nil, nil
	} else if err != nil {
		fmt.Println("One reading buff", err.Error())
		return packet, nil, err
	}
	buf = buf[:len]
	json.Unmarshal(buf, &packet)
	return packet, &addr, nil
}

func CreateServer(port string, handler func(net.PacketConn, *net.Addr, Packet)) *sync.WaitGroup {
	var wg sync.WaitGroup
	once.Do(func() {
		udpServer, err := net.ListenPacket("udp", ":"+port)
		if err != nil {
			log.Fatal(err)
		}
		instance = &udpServer

		//wait for the first connection
		fmt.Println("Waiting for the first connection...")
		packet, addr, err := readFromSocket()
		if err != nil {
			log.Fatal("Error when reading on the first connection:", err.Error())
		}
		serverPacketSystemHandler(packet.Data, addr)
		if err != nil {
			log.Fatal("error on json:", err.Error())
		}

		wg.Add(1)
		go listener(&wg, handler)
	})
	return &wg
}

func SendToAllClient(str, packetType string) {
	if instance == nil {
		log.Fatal("instance of UDP server is null")
		return
	}
	toSend := Packet{
		Type: packetType,
		Data: str,
	}

	res, err := json.Marshal(toSend)
	if err != nil {
		log.Fatal("error on json:", err.Error())
	}
	for _, addr := range addrs {
		(*instance).WriteTo(res, *addr)
	}
}

func SendToClient(str, packetType string, index int) {
	if instance == nil {
		log.Fatal("instance of UDP server is null")
		return
	}
	toSend := Packet{
		Type: packetType,
		Data: str,
	}
	res, err := json.Marshal(toSend)
	if err != nil {
		log.Fatal("error on json:", err.Error())
	}
	(*instance).WriteTo(res, *(addrs[index]))
}

func GetAllClientInfo() []string {
	var list []string
	for _, addr := range addrs {
		list = append(list, (*addr).String())
	}
	return list
}

func CloseServer() {
	SendToAllClient("close", "system")
	isServerClose = true
	defer (*instance).Close()
}
