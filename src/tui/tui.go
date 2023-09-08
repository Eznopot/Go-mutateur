package tui

import (
	"sync"

	"github.com/rivo/tview"
)

var app *tview.Application
var flex *tview.Flex
var once sync.Once

func Init() *tview.Application {
	once.Do(func() {
		app = tview.NewApplication()
		flex = tview.NewFlex()
	})
	return app
}

func Run() {
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

func createSelect(prompt string, options []string, handler func(int, string, string, rune)) *tview.List {
	listOptions := tview.NewList()
	listOptions.ShowSecondaryText(false)
	for index, option := range options {
		listOptions.AddItem(option, "", rune(49+index), nil)
	}

	listOptions.SetSelectedFunc(handler)
	return listOptions
}

func TuiCreateView(prompt string, options []string, handler func(int, string, string, rune)) {
	flexMenu := tview.NewFlex()
	flexMenu.Box.SetBorder(true).SetTitle("Menu")
	flexMenu.AddItem(createSelect(prompt, options, handler), 0, 1, true)

	boxPrincipal := tview.NewBox().SetBorder(true).SetTitle("Screen")

	flex.AddItem(flexMenu, 0, 1, true)
	flex.AddItem(boxPrincipal, 0, 3, false)
}

func Close() {
	app.Stop()
}
