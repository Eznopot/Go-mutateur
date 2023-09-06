package core

import (
	"go_mutateur/src/config"
	"go_mutateur/src/listener"
	"go_mutateur/src/udp"
)

func CoreClient() {
	udp.CreateConnection(config.GetConfig().Client.Address, config.GetConfig().Client.Port)
	udp.Receive(nil, func(bytes []byte) {
		listener.Do(string(bytes))
	})
}
