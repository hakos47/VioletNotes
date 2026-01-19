package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

// Purple shades
var (
	purplePrimary = color.RGBA{R: 128, G: 0, B: 128, A: 255}
	purpleFocus   = color.RGBA{R: 160, G: 32, B: 240, A: 255}
	purpleDark    = color.RGBA{R: 30, G: 0, B: 30, A: 255} // Very dark purple for background
)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNamePrimary:
		return purplePrimary
	case theme.ColorNameFocus:
		return purpleFocus
	case theme.ColorNameSelection:
		return purpleFocus
	// Optional: Tint the background slightly purple
	// case theme.ColorNameBackground:
	// 	 return purpleDark
	}

	// Fallback to standard dark theme for everything else
	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
