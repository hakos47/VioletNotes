package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

// VioletNotes Redesign Palette
var (
	// Obsidian Black
	colorBackground = color.RGBA{R: 10, G: 10, B: 10, A: 255}
	
	// Electric Violet
	colorPrimary = color.RGBA{R: 139, G: 92, B: 246, A: 255}
	
	// Deep Amethyst
	colorFocus = color.RGBA{R: 109, G: 40, B: 217, A: 255}
	
	// Transparent Violet for Glassmorphism effect
	colorSurface = color.RGBA{R: 30, G: 30, B: 35, A: 180}
)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return colorBackground
	case theme.ColorNamePrimary:
		return colorPrimary
	case theme.ColorNameFocus:
		return colorFocus
	case theme.ColorNameSelection:
		return colorFocus
	case theme.ColorNameInputBackground:
		return color.RGBA{R: 20, G: 20, B: 25, A: 255}
	case theme.ColorNameButton:
		return color.RGBA{R: 45, G: 45, B: 50, A: 255}
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNamePadding {
		return 8
	}
	if name == theme.SizeNameInputRadius {
		return 12
	}
	return theme.DefaultTheme().Size(name)
}