package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Onclik make title
func menu(app *tview.Application) *tview.List {
	list := tview.NewList().
		AddItem("Add New Todo", "", '1', nil).
		AddItem("Show Todo's", "", '2', nil)
	return list
}
func main() {

	app := tview.NewApplication()
	
	list := menu(app)
	list.SetBorder(true)
	flex := tview.NewFlex()

	flex.
		AddItem(list, 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true), 0, 3, false)

	transparentColor := tcell.NewHexColor(0x00000000) 

	flex.SetTitle(" Todo List Application ").SetBorder(true).SetBorderColor(tcell.ColorReset)
	flex.SetBorderColor(transparentColor)

	app.SetRoot(flex, true).SetFocus(flex).Run()
}
