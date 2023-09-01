package listener

import (
	"sync"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/keyboard"
)

var once sync.Once
var robot *gobot.Robot

func createKeyboardInstance(handler func(interface{})) {
	once.Do(func() {
		driver := keyboard.NewDriver()
		worker := func() {
			err := driver.On(keyboard.Key, func(data interface{}) {
				handler(data)
			})
			if err != nil {
				println(err.Error())
			}
		}
		robot = gobot.NewRobot("keyboard",
			[]gobot.Connection{},
			[]gobot.Device{driver},
			worker,
		)
	})
}

func GetKeyboardEvent(handler func(interface{})) {
	if robot == nil {
		println("iciii")
		createKeyboardInstance(handler)
	}
	robot.Start()
}
