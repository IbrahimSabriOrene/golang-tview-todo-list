package main

import (
	"github.com/rivo/tview"
)

type todo struct {
	Name, Description string
}


// Add navbar ->  <F1> menu, <F2> todo list(if exists), <F3> add todo, <F4> remove todo, <F5> update todo

var app =  tview.NewApplication()
var pages = tview.NewPages()
var list = tview.NewList()
var flex = tview.NewFlex()
var box = tview.NewBox()
var text = tview.NewTextView()
	
func showTodo(){
	pages.AddPage("Todo", text, true, true)
	pages.SwitchToPage("Todo")
	text.SetText(`Do something`)
} 
func showMenuTui() *tview.Pages{
	pages.AddPage("List", list, true, true)
	pages.SwitchToPage("List")
	list.AddItem("Show Todo's'","",'a',func() {
		showTodo()
	})
	list.AddItem("quit", "", 'q', func() {
		app.Stop()
	})
	return pages
}

func keywordBar(){
	
}

func main() {
		
	// List -> Title, Description
	// ------------ V1 -------------
	// Show contents
	// Add List
	// Remove List
	// Update List

	// ------------ V2 -------------
	// Add checkbox
	// The project's cookie resets every 24 hour - changeable
	menu := showMenuTui()
	app.SetRoot(menu, true).SetFocus(menu).Run()
}
