package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/valli0x/ens-sig/boxes"
)

func main() {
	a := app.New()
	w := a.NewWindow("Multi-Sign App app")
	w.Resize(fyne.NewSize(600, 600))
	w.CenterOnScreen()

	root, err := boxes.Root(w, boxes.SignContainer, boxes.CheckContainer)
	if err != nil {
		panic(err)
	}

	w.SetContent(root)
	w.ShowAndRun()
}
