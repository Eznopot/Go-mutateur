package udp

import (
	"fmt"
	"log"
	"net"
	"sync"
)

var instance *net.PacketConn
var addrs []*net.Addr

var once sync.Once

func listener(wg *sync.WaitGroup, handler func(net.PacketConn, net.Addr, []byte)) {
	defer wg.Done()
	for {
		buf := make([]byte, 1024)
		_, addr, err := (*instance).ReadFrom(buf)
		addrs = (append(addrs, &addr))
		fmt.Println("New client connected on:", addr)
		if err != nil {
			break
		}
		handler(*instance, addr, buf)
	}
}

func CreateServer(port string, handler func(net.PacketConn, net.Addr, []byte)) (*net.PacketConn, *sync.WaitGroup) {
	var wg sync.WaitGroup
	once.Do(func() {
		udpServer, err := net.ListenPacket("udp", ":"+port)
		if err != nil {
			log.Fatal(err)
		}
		instance = &udpServer

		//wait for the first connection
		fmt.Println("Waiting for the first connection...")
		buf := make([]byte, 1024)
		_, addr, err := (*instance).ReadFrom(buf)
		var tmp []*net.Addr
		tmp = append(tmp, &addr)
		addrs = tmp

		fmt.Println("First client connected:", addr)
		if err != nil {
			fmt.Println("error here")
		}

		wg.Add(1)
		go listener(&wg, handler)
	})
	return instance, &wg
}

func SendToAllClient(str string) {
	if instance == nil {
		log.Fatal("instance of UDP server is null")
		return
	}
	for _, addr := range addrs {
		(*instance).WriteTo([]byte(str), *addr)
	}
}

func SendToClient(str string, index int) {
	if instance == nil {
		log.Fatal("instance of UDP server is null")
		return
	}
	(*instance).WriteTo([]byte(str), *(addrs[index]))
}

func GetAllClientInfo() []string {
	var list []string
	for _, addr := range addrs {
		list = append(list, (*addr).String())
	}
	return list
}
