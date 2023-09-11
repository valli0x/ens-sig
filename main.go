package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/valli0x/ens-sig/containers"
)

func main() {
	a := app.New()
	w := a.NewWindow("ENS signing app")
	w.Resize(fyne.NewSize(600, 500))
	w.CenterOnScreen()

	rootContainer, err := containers.Root(w, containers.SignContainer, containers.CheckContainer)
	if err != nil {
		panic(err)
	}

	w.SetContent(rootContainer)
	w.ShowAndRun()
}
