package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// CommandEntry is a custom Entry widget that intercepts specific shortcuts
// (like Ctrl+S) to trigger application-level actions even when focused.
type CommandEntry struct {
	widget.Entry
	onSave func()
}

// NewCommandEntry creates a new CommandEntry with a save callback.
func NewCommandEntry(onSave func()) *CommandEntry {
	entry := &CommandEntry{onSave: onSave}
	entry.ExtendBaseWidget(entry)
	return entry
}

// TypedShortcut intercepts shortcuts. It specifically looks for Ctrl+S
// to trigger the onSave callback. All other shortcuts are delegated to the base Entry.
func (e *CommandEntry) TypedShortcut(s fyne.Shortcut) {
	if custom, ok := s.(*desktop.CustomShortcut); ok {
		if custom.KeyName == fyne.KeyS && custom.Modifier == fyne.KeyModifierControl {
			if e.onSave != nil {
				e.onSave()
				return
			}
		}
	}
	e.Entry.TypedShortcut(s)
}
