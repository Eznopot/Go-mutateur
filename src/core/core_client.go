package core

import (
	"go_mutateur/src/config"
	"go_mutateur/src/listener"
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
	udp_client.CreateConnection(config.GetConfig().Client.Address, config.GetConfig().Client.Port)
	udp_client.Receive(nil, func(packet udp.Packet) {
		listener.Do(packet.Data)
	})
}
