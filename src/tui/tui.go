package tui

import (
	"sync"

	"github.com/rivo/tview"
)

var app *tview.Application
var flex *tview.Flex
var once sync.Once
var textViewGlobal *tview.TextView
var listOptions *tview.List

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

// The FocusMenu function sets the focus on the first item in a flex container.
func FocusMenu() {
	app.SetFocus(flex.GetItem(0))
}

// The function `FocusScreen` sets the focus on the second item in a flex container.
func FocusScreen() {
	app.SetFocus(flex.GetItem(1))
}

// The function creates a selectable list with options and shortcuts, and calls a handler function when
// an option is selected.
//
// Args:
//   prompt (string): The prompt is a string that will be displayed at the top of the select menu,
// prompting the user to make a selection.
//   options ([]string): A slice of strings representing the options to be displayed in the select
// list.
//   shortcuts ([]rune): The `shortcuts` parameter is a slice of `rune` values that represents the
// keyboard shortcuts for each option in the `options` slice. Each shortcut corresponds to an option in
// the `options` slice at the same index. For example, if the `options` slice has three elements and
//   handler: The `handler` parameter is a function that takes four arguments: int, string, string, rune
func createSelect(prompt string, options []string, shortcuts []rune, handler func(int, string, string, rune)) *tview.List {
	listOptions = tview.NewList()
	listOptions.ShowSecondaryText(false)
	AddMenuOption(options, shortcuts)

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
// user before they make a selection from the menu. It serves as a prompt or instruction for the user
// to follow.
//   options ([]string): The `options` parameter is a slice of strings that represents the menu options
// that will be displayed to the user. Each string in the slice represents an option in the menu.
//   shortcuts ([]rune): The `shortcuts` parameter is a slice of `rune` values that represents the
// keyboard shortcuts for each option in the menu. Each `rune` value corresponds to a specific key on
// the keyboard. For example, if you want to assign the shortcut "A" to the first option,
//   handler: The `handler` parameter is a function that will be called when an option is selected in
// the menu. It takes four arguments: int, string, string, rune
func TuiCreateView(prompt string, options []string, shortcuts []rune, handler func(int, string, string, rune)) {
	flexMenu := tview.NewFlex()
	flexMenu.Box.SetBorder(true).SetTitle("Menu")
	flexMenu.AddItem(createSelect(prompt, options, shortcuts, handler), 0, 1, true)

	boxPrincipal := tview.NewFlex()
	boxPrincipal.SetBorder(true).SetTitle("Screen")
	createText()
	boxPrincipal.AddItem(textViewGlobal, 0, 1, false)

	flex.AddItem(flexMenu, 0, 1, true)
	flex.AddItem(boxPrincipal, 0, 4, false)
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

// The function removes a menu option at a specified index.
//
// Args:
//   index (int): The index parameter is an integer that represents the position of the menu option
// that you want to remove. It is used to identify the specific menu option that you want to remove
// from the list.
func RemoveMenuOption(index int) {
	listOptions.RemoveItem(index)
}

// The function `RemoveMenuOptionByString` removes a menu option from a list based on a given string
// and returns the index of the removed item.
//
// Args:
//   str (string): The parameter "str" is a string that represents the menu option that needs to be
// removed from the listOptions.
//
// Returns:
//   an integer value. If the menu option specified by the input string is found and successfully
// removed, the function returns the index of the removed item. If the menu option is not found, the
// function returns -1.
func RemoveMenuOptionByString(str string) int {
	size := listOptions.GetItemCount()
	for i := 0; i < size; i++ {
		main, _ := listOptions.GetItemText(i)
		if main == str {
			listOptions.RemoveItem(i)
			return i
		}
	}
	return -1
}

// The function "AddMenuOption" adds menu options with corresponding shortcuts to a list.
//
// Args:
//   options ([]string): A slice of strings representing the menu options. Each string represents a
// different option in the menu.
//   shortcuts ([]rune): The `shortcuts` parameter is a slice of `rune` types. Each `rune` represents a
// keyboard shortcut for a menu option.
func AddMenuOption(options []string, shortcuts []rune) {
	for index, option := range options {
		listOptions.AddItem(option, "", shortcuts[index], nil)
	}
}

// The function InsertMenuOption inserts a menu option at a specified index in a list.
//
// Args:
//   option (string): The option parameter is a string that represents the menu option that you want to
// insert.
//   index (int): The index parameter is an integer that represents the position at which the menu
// option should be inserted.
func InsertMenuOption(option string, index int) {
	listOptions.InsertItem(index, option, "", rune(49+index), nil)
}

// The Close function stops the application.
func Close() {
	app.Stop()
}
