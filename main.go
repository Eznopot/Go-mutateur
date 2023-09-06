package main

import (
	"fmt"
	"go_mutateur/src/config"
	"go_mutateur/src/core"
	"os"
)

func main() {
	args := os.Args
	config.GetConfig()
	if len(args) == 2 && args[1] == "server" {
		println("server mode launch...")
		core.CoreServer()
		return
	} else if len(args) == 2 && args[1] == "client" {
		println("client mode launch....")
		core.CoreClient()
		return
	}
	fmt.Println("use argument to select server or client version")
}
