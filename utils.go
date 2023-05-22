package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func NewSpacer(width float32, height float32) *canvas.Rectangle {
	w := canvas.NewRectangle(&color.NRGBA{0, 0, 0, 0})
	w.SetMinSize(fyne.NewSize(width, height))
	return w
}
