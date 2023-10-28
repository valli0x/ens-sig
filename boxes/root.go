package boxes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	backName = "back"
	desc     = `
Description
This is a program for verifying the ECDSA signature using the Ethereum blockchain. 
It gets a public key from the hash of the file and the signature. 
Next, it receives a hash from the public key or an address in Ethereum. 
Next, it accesses the Ethereum blockchain through the client 
and receives a domain (or name) by reverse resolving. 
Compares the entered name and the received one and outputs the result 
of the comparison`
)

type BoxConstructor func(w fyne.Window, back *widget.Button) (name string, _ *fyne.Container, _ error)

func Root(w fyne.Window, boxList ...BoxConstructor) (*fyne.Container, error) {
	root := container.NewVBox()
	back := widget.NewButton(backName, func() {
		w.SetContent(root)
	})

	head := widget.NewLabelWithStyle("Main menu", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	root.Add(head)

	description := widget.NewLabelWithStyle(desc, fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: false})
	root.Add(description)

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
