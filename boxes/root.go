package boxes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	backName = "back"
)

type BoxConstructor func(w fyne.Window, back *widget.Button) (name string, _ *fyne.Container, _ error)

func Root(w fyne.Window, boxList ...BoxConstructor) (*fyne.Container, error) {
	root := container.NewVBox()
	back := widget.NewButton(backName, func() {
		w.SetContent(root)
	})

	head := widget.NewLabelWithStyle("Main menu", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	root.Add(head)
	root.Add(layout.NewSpacer())

	for _, constructor := range boxList {
		name, containerList, err := constructor(w, back)
		if err != nil {
			return nil, err
		}

		button := widget.NewButton(name, func() {
			w.SetContent(containerList)
		})
		root.Add(button)
	}

	return root, nil
}
