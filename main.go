package main

import (
	"fmt"
	"go_mutateur/src/core"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 2 && args[1] == "server" {
		println("server mode")
		core.CoreServer()
		return
	} else if len(args) == 2 && args[1] == "client" {
		println("client mode")
		core.CoreClient()
		return
	}
	fmt.Println("use argument to select server or client version")
}
