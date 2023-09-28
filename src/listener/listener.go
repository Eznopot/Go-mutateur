package listener

import (
	"encoding/json"
	"fmt"
	"go_mutateur/src/config"
	"go_mutateur/src/tui"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

// The `Event` function takes three function parameters, `onStart`, `onEnd`, and `handler`, and sets up
// an event loop that listens for keyboard events and calls the appropriate functions based on the
// events received.
//
// Args:
//
//	`onStart`: The onStart parameter is a function that will be called at the beginning of the Event
//
// function. It is typically used to set up any necessary resources or perform any initialization tasks
// before the event loop starts.
//
//	onEnd: The `onEnd` parameter is a function that will be called when the event loop ends. It is
//
// typically used to clean up any resources or perform any necessary final actions.
//
//	handler: The `handler` parameter is a function that takes a single argument of type `hook.Event`.
//
// This function is responsible for handling the events received from the `evChan` channel.
func Event(onStart, onEnd func(), handler func(hook.Event)) {
	evChan := hook.Start()
	onStart()
	defer hook.End()

	for ev := range evChan {
		if ev.Kind == hook.KeyDown && ev.Keychar == 27 {
			tui.AddLog("---------------------------------------")
			onEnd()
			return
		}
		handler(ev)
	}
	onEnd()
}

// The function Do takes in an event string, parses it into a `hook.Event` struct, and performs
// different actions based on the kind of event.
//
// Args:
//
//	eventString (string): The `eventString` parameter is a string that represents a JSON object (hook.Event)
//
// containing information about an event. This string is used to deserialize the event object using the
// `json.Unmarshal` function.
func Do(eventString string) {
	var event hook.Event
	err := json.Unmarshal([]byte(eventString), &event)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	if event.Kind == hook.KeyDown {
		robotgo.KeyDown(hook.RawcodetoKeychar(event.Rawcode))
	} else if event.Kind == hook.KeyUp {
		robotgo.KeyUp(hook.RawcodetoKeychar(event.Rawcode))
	} else if event.Kind == hook.MouseDown {
		var clickAction string
		if event.Button == 1 {
			clickAction = "left"
		} else if event.Button == 2 {
			clickAction = "right"
		}
		robotgo.Click(clickAction)
	} else if event.Kind == hook.MouseWheel {
		if event.Rotation > 0 {
			robotgo.ScrollMouse(config.GetConfig().Config.ScrollSpeed, "down")
		} else {
			robotgo.ScrollMouse(-config.GetConfig().Config.ScrollSpeed, "up")
		}
	} else if event.Kind == hook.MouseDrag {
		robotgo.DragSmooth(int(event.X), int(event.Y), 0.1, 0.3)
	} else if event.Kind == hook.MouseMove {
		if config.GetConfig().Config.SmoothMode {
			robotgo.MoveSmooth(int(event.X), int(event.Y), 0.1, 0.3, config.GetConfig().Config.SmoothDelay)
		} else {
			robotgo.Move(int(event.X), int(event.Y))
		}
	}
}
