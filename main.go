package main

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// App represents the main application state and UI references.
type App struct {
	fyneApp fyne.App
	window  fyne.Window

	// Data
	notes            []Note
	currentNoteIndex int

	// Bindings
	titleBind   binding.String
	contentBind binding.String

	// UI Components
	list         *widget.List
	titleEntry   *CommandEntry
	contentEntry *CommandEntry
	btnSave      *widget.Button
	btnDelete    *widget.Button
	btnAdd       *widget.Button
}

func main() {
	// Force Fyne language before any app initialization
	os.Setenv("FYNE_LANG", "en-US")

	myApp := &App{
		fyneApp:          app.New(),
		currentNoteIndex: -1,
		titleBind:        binding.NewString(),
		contentBind:      binding.NewString(),
	}

	// Apply Custom Theme
	myApp.fyneApp.Settings().SetTheme(&myTheme{})

	// Window Setup
	myApp.window = myApp.fyneApp.NewWindow("VioletNotes")
	myApp.window.Resize(fyne.NewSize(900, 600))

	// Load Data
	var err error
	myApp.notes, err = LoadNotes()
	if err != nil {
		fmt.Printf("Error loading notes: %v\n", err)
		myApp.notes = []Note{}
	}

	// Initialize UI
	myApp.buildUI()

	// Global Shortcuts
	myApp.window.Canvas().AddShortcut(
		&desktop.CustomShortcut{KeyName: fyne.KeyS, Modifier: fyne.KeyModifierControl},
		func(shortcut fyne.Shortcut) { myApp.saveCurrentNote() },
	)

	myApp.window.ShowAndRun()
}

func (a *App) buildUI() {
	// --- List View ---
	a.list = widget.NewList(
		func() int { return len(a.notes) },
		func() fyne.CanvasObject {
			label := widget.NewLabel("Note Title")
			label.Truncation = fyne.TextTruncateEllipsis
			return label
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(a.notes[i].Title)
		},
	)

	a.list.OnSelected = a.onNoteSelected
	a.list.OnUnselected = a.onNoteUnselected

	// --- Editor View ---
	a.titleEntry = NewCommandEntry(a.saveCurrentNote)
	a.titleEntry.Bind(a.titleBind)
	a.titleEntry.PlaceHolder = "Note Title"

	a.contentEntry = NewCommandEntry(a.saveCurrentNote)
	a.contentEntry.Bind(a.contentBind)
	a.contentEntry.PlaceHolder = "Start typing your note..."
	a.contentEntry.Wrapping = fyne.TextWrapWord
	a.contentEntry.MultiLine = true

	// --- Actions ---
	a.btnSave = widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), a.saveCurrentNote)
	a.btnDelete = widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), a.deleteCurrentNote)
	a.btnAdd = widget.NewButtonWithIcon("New Note", theme.ContentAddIcon(), a.addNewNote)

	// Initial State (Disabled until selection)
	a.setEditorEnabled(false)

	// --- Layout Composition ---
	
	// Left Pane: Add Button + List
	leftPane := container.NewBorder(
		container.NewPadded(a.btnAdd),
		nil, nil, nil,
		a.list,
	)

	// Right Pane: Title + Toolbar + Content
	editorToolbar := container.NewHBox(layout.NewSpacer(), a.btnDelete, a.btnSave)
	rightPane := container.NewBorder(
		container.NewVBox(a.titleEntry, editorToolbar),
		nil, nil, nil,
		a.contentEntry,
	)

	// Split Container
	split := container.NewHSplit(leftPane, container.NewPadded(rightPane))
	split.SetOffset(0.3)

	// Background Gradient
	gradient := canvas.NewLinearGradient(
		color.RGBA{R: 45, G: 0, B: 60, A: 255}, // Deep Purple
		color.RGBA{R: 0, G: 0, B: 0, A: 255},   // Black
		0,
	)

	// Final Layout
	a.window.SetContent(container.NewMax(gradient, split))
}

// Logic Methods

func (a *App) saveCurrentNote() {
	if a.currentNoteIndex < 0 || a.currentNoteIndex >= len(a.notes) {
		return
	}

	// Sync binding to data model
	title, _ := a.titleBind.Get()
	content, _ := a.contentBind.Get()

	a.notes[a.currentNoteIndex].Title = title
	a.notes[a.currentNoteIndex].Content = content

	if err := SaveNotes(a.notes); err != nil {
		fmt.Printf("Error saving notes: %v\n", err)
		// Ideally show a dialog here
	}
	a.list.Refresh()
}

func (a *App) deleteCurrentNote() {
	if a.currentNoteIndex < 0 || a.currentNoteIndex >= len(a.notes) {
		return
	}

	// Remove from slice
	a.notes = append(a.notes[:a.currentNoteIndex], a.notes[a.currentNoteIndex+1:]...)
	
	if err := SaveNotes(a.notes); err != nil {
		fmt.Printf("Error saving after delete: %v\n", err)
	}

	a.list.UnselectAll()
	a.list.Refresh()
}

func (a *App) addNewNote() {
	newNote := Note{
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		Title:     "Untitled Note",
		Content:   "",
		CreatedAt: time.Now(),
	}

	// Prepend to list
	a.notes = append([]Note{newNote}, a.notes...)
	
	if err := SaveNotes(a.notes); err != nil {
		fmt.Printf("Error saving new note: %v\n", err)
	}

	a.list.Refresh()
	a.list.Select(0) // Select the newly created note
}

func (a *App) onNoteSelected(id widget.ListItemID) {
	a.currentNoteIndex = id
	if id >= 0 && id < len(a.notes) {
		a.titleBind.Set(a.notes[id].Title)
		a.contentBind.Set(a.notes[id].Content)
		a.setEditorEnabled(true)
	}
}

func (a *App) onNoteUnselected(id widget.ListItemID) {
	a.currentNoteIndex = -1
	a.titleBind.Set("")
	a.contentBind.Set("")
	a.setEditorEnabled(false)
}

func (a *App) setEditorEnabled(enabled bool) {
	if enabled {
		a.titleEntry.Enable()
		a.contentEntry.Enable()
		a.btnSave.Enable()
		a.btnDelete.Enable()
	} else {
		a.titleEntry.Disable()
		a.contentEntry.Disable()
		a.btnSave.Disable()
		a.btnDelete.Disable()
	}
}