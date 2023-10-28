package boxes

import (
	"encoding/hex"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/tyler-smith/go-bip39"
	"github.com/valli0x/ens-sig/filehash"
	"github.com/valli0x/ens-sig/signfile"
)

const (
	signNameFunc = "sign"
)

func SignContainer(w fyne.Window, back *widget.Button) (name string, _ *fyne.Container, _ error) {
	head := widget.NewLabelWithStyle("Sign", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	privLb := widget.NewLabelWithStyle("Private Key or Mnemonic", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	privEntry := widget.NewPasswordEntry()

	filePath := widget.NewLabelWithStyle("Fite path", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	filePathEntry := widget.NewEntry()

	sigLb := widget.NewLabelWithStyle("Signature", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	sigStr := ""
	sigStrBind := binding.BindString(&sigStr)
	sigEntry := widget.NewEntryWithData(sigStrBind)

	errStr := ""
	errStrBind := binding.BindString(&errStr)
	errEntry := widget.NewLabelWithData(errStrBind)

	signBtn := widget.NewButton("sign", func() {
		errStrBind.Set("")
		errStrBind.Reload()

		filehash, err := filehash.FileHash(filePathEntry.Text)
		if err != nil {
			errStrBind.Set("error: " + err.Error())
			errStrBind.Reload()
			return
		}

		if bip39.IsMnemonicValid(privEntry.Text) {
			privByte, err := bip39.EntropyFromMnemonic(privEntry.Text)
			if err != nil {
				errStrBind.Set("error: " + err.Error())
				errStrBind.Reload()
				return
			}
			privEntry.Text = fmt.Sprint(privByte)
			memzero(privByte)
		}

		signature, err := signfile.SignHash(filehash, privEntry.Text)
		if err != nil {
			errStrBind.Set("error: " + err.Error())
			errStrBind.Reload()
			return
		}

		sigStrBind.Set(hex.EncodeToString(signature))
		sigStrBind.Reload()

		privEntry.Text = ""
	})

	signBox := container.NewVBox(
		head,
		privLb,
		privEntry,
		filePath,
		filePathEntry,
		sigLb,
		sigEntry,
		layout.NewSpacer(),
		errEntry,
		signBtn,
		back,
	)

	return signNameFunc, signBox, nil
}

func memzero(b []byte) {
	if b == nil {
		return
	}
	for i := range b {
		b[i] = 0
	}
}
