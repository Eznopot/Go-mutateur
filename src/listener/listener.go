package listener

import (
	"encoding/json"
	"fmt"

	hook "github.com/robotn/gohook"
)

func Event(handler func(hook.Event)) {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		if ev.Kind == hook.KeyDown && ev.Keychar == 27 {
			return
		}
		handler(ev)
	}
}

func Do(eventString string) {
	var event hook.Event
	err := json.Unmarshal([]byte(eventString), &event)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("keyboard event:", event)
}
