package core

import (
	"go_mutateur/src/listener"
	"go_mutateur/src/udp"
)

func CoreClient() {
	//connect to the server
	udp.CreateConnection()
	//receive instruction
	udp.Receive(nil, func(bytes []byte) {
		listener.Do(string(bytes))
	})
}
