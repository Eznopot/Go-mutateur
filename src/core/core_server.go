package core

import (
	"fmt"
	"go_mutateur/src/listener"
	"net"
	"syscall"

	"gobot.io/x/gobot/platforms/keyboard"
)

func handlerServer(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	fmt.Println("recus:", string(buf))
}

func keyboardEventHandler(data interface{}) {
	key := data.(keyboard.KeyEvent)
	if key.Key == keyboard.Escape {
		//kill robot
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	} else {
		fmt.Println("keyboard event!", key, key.Char)
	}
}

func CoreServer() {
	//First wait for client connection
	//udp.CreateServer(handlerServer)
	listener.GetKeyboardEvent(keyboardEventHandler) //dont work if relauch do everywith in double
	//in a for with condition: if response in prompt is exit exit loop and programme
	//when first client is connected display prompt to select wich computer need to be controlled
	//make for loop with keyboard event and mouse event to get keyboard input
	//if escape is pressed exit loop
	//send input to selected client
	//close udp server
}
