package main

import (
	"fmt"
	"io"
	"os"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)



// Add navbar ->  <F1> menu, <F2> todo list(if exists), <F3> add todo, <F4> remove todo, <F5> update todo

var app = tview.NewApplication()
var pages = tview.NewPages()
var navbar = tview.NewList()
var list = tview.NewList()
var flex = tview.NewFlex()
var box = tview.NewBox()
var text = tview.NewTextArea()
var content = tview.NewTextView()


func addTodo() {

	pages.AddPage("AddTodo", text, true, true)
	pages.SwitchToPage("AddTodo")
}

func showTodo(){

	pages.AddPage("ShowTodo", content, true, true)
	pages.SwitchToPage("ShowTodo")
	
}

func navbarTui() *tview.Pages {

	pages.AddPage("List", list, true, true)
	list.Clear()
	list.SetTitle("Menu")

	list.SetBorder(true)
	list.AddItem("Add Todo's'", "", 'a', func() {
		addTodo()
	})
	
	list.AddItem("quit", "", 'q', func() {
		app.Stop()
	})
	pages.SwitchToPage("List")
	return pages
}
func saveTodo(app *tview.Application) {
	filename := "Todo.txt"
	storedText := text.GetText()
   // Add the line number and a space before the line
        fmt.Fprintf(tempFile, "%d %s\n", lineNumber, line)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating the file")
		return
	}
	defer file.Close()

	_, err = file.WriteString(storedText)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while writing file")
		return
	}
	text.SetText(storedText, true)
	content.SetText(storedText)

}
func loadTodo(app *tview.Application, textarea *tview.TextView) {
	//on first start this should've start
	filename:="Todo.txt"
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		return
	}

	textarea.SetText(string(data))
	text.SetText(string(data),true)
}

func handleKeyPress(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyF1:
		pages.SwitchToPage("List")
		app.SetFocus(list)
	case tcell.KeyF2:
		app.SetFocus(content)
	case tcell.KeyCtrlS:
		saveTodo(app)
		
	}
	return event
}

func main() {

	// Write stuff,
	//Press <C-s>
	//On save, project saves in a text file the output will be like this->
	// (1) DO SOMETHING21311
	// (2) DO ANOTHERTHING 313211
	// (3) DO GOOD THING
	// (4) GREAT stuff
	// On pressing enter, remove the current textarea's text, add it to the textview. Example <Enter Pressed> (number) Keep learning golang
	// To remove anything, we are gonna use Id = 1 etc. but instead ID = 1, we need better way to remove it, prefer line removal
	// update by query???
	// Later we are gonna connect MySql Database

	// The project's cookie resets every 24 hour - changeable
	//
	loadTodo(app,content)
	menu := navbarTui()
	flex.AddItem(menu, 0, 1, false)
	//flex.AddItem(tview.NewTextView(), 0,1, false)
	flex.AddItem( content,0, 3, false)
	app.SetInputCapture(handleKeyPress)
	app.SetRoot(flex, true).SetFocus(menu).Run()
}
