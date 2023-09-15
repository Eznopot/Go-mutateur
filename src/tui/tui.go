package tui

import (
	"sync"

	"github.com/rivo/tview"
)

var app *tview.Application
var flex *tview.Flex
var once sync.Once
var textViewGlobal *tview.TextView

// The `Init` function initializes a new `tview.Application` and returns it. (Actually a Singleton)
func Init() *tview.Application {
	once.Do(func() {
		app = tview.NewApplication()
		flex = tview.NewFlex()
	})
	return app
}

// The Run function sets the root of the application and runs it, panicking if there is an error.
func Run() {
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

// The function "createSelect" creates a selectable list with options and a prompt.
//
// Args:
//   prompt (string): A string that represents the prompt or question to be displayed before the list
// of options.
//   options ([]string): The `options` parameter is a slice of strings that represents the list of
// options to be displayed in the select list. Each string in the slice represents an option that the
// user can choose from.
//   handler: The `handler` parameter is a function that takes four arguments: int, string, string, rune
func createSelect(prompt string, options []string, handler func(int, string, string, rune)) *tview.List {
	listOptions := tview.NewList()
	listOptions.ShowSecondaryText(false)
	for index, option := range options {
		listOptions.AddItem(option, "", rune(49+index), nil)
	}

	listOptions.SetSelectedFunc(handler)
	return listOptions
}

// The function `createText()` returns a new `tview.TextView` object with certain settings applied.
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

// The function `TuiCreateView` creates a TUI (Text User Interface) view with a menu and a screen.
//
// Args:
//   prompt (string): The prompt is a string that represents the message or question displayed to the
// user in the menu. It prompts the user to make a selection or input some information.
//   options ([]string): The `options` parameter is a slice of strings that represents the available
// options for the menu. Each string in the slice represents an option that the user can select from
// the menu.
//   handler: The `handler` parameter is a function that takes four arguments: int, string, string, rune
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

// The AddLog function appends a string to a log array, removes the oldest log if the array exceeds 20
// elements, and updates a text view with the log messages.
//
// Args:
//   str (string): The parameter "str" is a string that represents the log message that needs to be
// added to the logs.
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

// The Close function stops the application.
func Close() {
	app.Stop()
}
