package main

import (
	"fmt"
	"go_mutateur/src/config"
	core_client "go_mutateur/src/core/client"
	core_server "go_mutateur/src/core/server"
	"os"
)

func main() {
	args := os.Args
	config.GetConfig()
	if len(args) == 2 && args[1] == "server" {
		println("server mode launch...")
		core_server.CoreServer()
		return
	} else if len(args) == 2 && args[1] == "client" {
		println("client mode launch....")
		core_client.CoreClient()
		return
	}
	fmt.Println("use argument to select server or client version")
}
