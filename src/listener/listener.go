package listener

import (
	"encoding/json"
	"fmt"

	"github.com/go-vgo/robotgo"
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
	if event.Kind == hook.KeyDown {
		robotgo.KeyTap(string(event.Keychar))
	} else if event.Kind == hook.MouseDown {
		var clickAction string
		if event.Button == 1 {
			clickAction = "left"
		} else if event.Button == 2 {
			clickAction = "right"
		}
		robotgo.Click(clickAction)
	} else if event.Kind == hook.MouseWheel {
		robotgo.Scroll(int(event.X), int(event.Y))
	} else if event.Kind == hook.MouseDrag {
		robotgo.DragSmooth(int(event.X), int(event.Y))
	} else if event.Kind == hook.MouseMove {
		robotgo.Move(int(event.X), int(event.Y))
	}
	fmt.Println("keyboard event:", event)
}