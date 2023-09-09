package tui

import (
	"sync"

	"github.com/rivo/tview"
)

var app *tview.Application
var flex *tview.Flex
var once sync.Once
var textViewGlobal *tview.TextView

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

func createText() *tview.TextView {
	textViewGlobal = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	return textViewGlobal
}

func TuiCreateView(prompt string, options []string, handler func(int, string, string, rune)) {
	flexMenu := tview.NewFlex()
	flexMenu.Box.SetBorder(true).SetTitle("Menu")
	flexMenu.AddItem(createSelect(prompt, options, handler), 0, 1, true)

	boxPrincipal := tview.NewFlex()
	boxPrincipal.SetBorder(true).SetTitle("Screen")
	createText()
	boxPrincipal.AddItem(textViewGlobal, 0, 1, false)

	flex.AddItem(flexMenu, 0, 1, true)
	flex.AddItem(boxPrincipal, 0, 1, false)
}

var logs []string

func AddLog(str string) {
	var logStr string

	if len(logs) > 20 {
		logs = logs[1:]
	}
	logs = append(logs, str)

	for i := range logs {
		logStr += logs[len(logs)-i-1] + "\n"
	}
	textViewGlobal.SetText(logStr)
}

func Close() {
	app.Stop()
}
