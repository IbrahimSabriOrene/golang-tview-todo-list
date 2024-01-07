package main

import (
	"bufio"
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
var inputField = tview.NewInputField()
var text = tview.NewTextArea()
var content = tview.NewTextView()
var keyinfo1 = tview.NewTextView()
var keyinfo2 = tview.NewTextView()
var keyinfo3 = tview.NewTextView()

func addTodo() {

	pages.AddPage("AddTodo", text, true, true)
	pages.SwitchToPage("AddTodo")
}

func showTodo() {

	pages.AddPage("ShowTodo", content, true, true)
	pages.SwitchToPage("ShowTodo")

}

func navbarTui() *tview.Pages {

	pages.AddPage("List", list, true, true)
	//	list.Clear()
	//	list.SetTitle("Menu")
	//
	//	list.SetBorder(true)
	//	list.AddItem("Add Todo's'", "", 'a', func() {
	//		addTodo()
	//	})
	//
	//	list.AddItem("quit", "", 'q', func() {
	//		app.Stop()
	//	})
	inputField.SetLabel("Task")

	pages.SwitchToPage("List")
	return pages
}
func saveTodo(app *tview.Application) {
	filename := "Todo.txt"
	storedText := text.GetText()
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating the file")
		return
	}
	defer file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while writing file")
		return
	}

	_, err = file.WriteString(storedText)
	// Add the line number and a space before the line
	setLineNumber() // Move this to somewhere else it keeps repeating
	text.SetText(storedText, true)
	content.SetText(storedText)

}
func loadTodo(app *tview.Application, textarea *tview.TextView) {
	//on first start this should've start
	filename := "Todo.txt"
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
	text.SetText(string(data), true)
}

func setLineNumber() {
	filename := "Todo.txt" // Replace with your filename

	// Open the file in read mode
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a new temporary file to store the modified content
	tempFile, err := os.Create("temp_file.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating temporary file: %v\n", err)
		return
	}
	defer tempFile.Close()

	// Iterate through each line of the original file
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		// Get the current line
		line := scanner.Text()

		// Add the line number and a space before the line
		fmt.Fprintf(tempFile, "%d %s\n", lineNumber, line)

		lineNumber++
	}

	// Handle any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning file: %v\n", err)
		return
	}

	// Close both files
	file.Close()
	tempFile.Close()

	// Overwrite the original file with the modified content
	if err := os.Rename("temp_file.txt", filename); err != nil {
		fmt.Fprintf(os.Stderr, "Error renaming temporary file: %v\n", err)
		return
	}

	fmt.Println("Numbers added successfully to each line of the file!")
}

func handleKeyPress(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyF1:
		pages.SwitchToPage("List")
		app.SetFocus(text)
	case tcell.KeyF3:
		app.SetFocus(content)
	case tcell.KeyCtrlS:
		saveTodo(app)
	case tcell.KeyF2:
		app.Stop()
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
	// On pressing enter, remove the current textarea's text, add it to the textview with (number)
	// To remove anything, we are gonna use Id = 1 etc. but instead ID = 1, we need better way to remove it, prefer line removal
	// update by query???
	// Later we are gonna connect MySql Database

	// The project's cookie resets every 24 hour - changeable
	//
	keyinfo1.SetText("<F1> Add Todo").SetTextAlign(tview.AlignCenter)
	keyinfo2.SetText("<Q> Quit").SetTextAlign(tview.AlignCenter)
	keyinfo3.SetText("<F3> Quit and Save").SetTextAlign(tview.AlignCenter)

	//menu := navbarTui()
	flex := tview.NewFlex()

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(flex.AddItem(keyinfo1, 0, 2, false).
			AddItem(keyinfo2, 0, 2, false).
			AddItem(keyinfo3, 0, 2, false), 2, 0, 1, 3, 1, 1, false).
		AddItem(inputField, 0, 0, 1, 3, 0, 0, false).
		AddItem(content, 1, 0, 1, 3, 0, 0, false)

	content.SetTextAlign(tview.AlignCenter)

	loadTodo(app, content)
	//flex.AddItem(menu, 0, 1, false)
	//flex.AddItem(tview.NewTextView(), 0,1, false)
	//flex.AddItem(content, 0, 3, false)
	app.SetInputCapture(handleKeyPress)
	app.SetRoot(grid, true).SetFocus(inputField).Run()
}
