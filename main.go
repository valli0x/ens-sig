package main

import (
	"encoding/hex"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tyler-smith/go-bip39"
	"github.com/valli0x/ens-sig/ens"
	"github.com/valli0x/ens-sig/filehash"
	"github.com/valli0x/ens-sig/signfile"
)

func main() {
	a := app.New()
	w := a.NewWindow("ENS signing app")
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()

	// signContainer, err := signContainer()
	// if err != nil {
	// 	panic(err)
	// }

	checkContainer, err := checkContainer()
	if err != nil {
		panic(err)
	}

	w.SetContent(checkContainer)
	w.ShowAndRun()
}

func signContainer() (*fyne.Container, error) {
	head := widget.NewLabelWithStyle("Sign", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	privLb := widget.NewLabelWithStyle("Private Key or Mnemonic", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	privEntry := widget.NewPasswordEntry()

	filePathLb := widget.NewLabelWithStyle("File path", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	filePathEntry := widget.NewEntry()

	sigLb := widget.NewLabelWithStyle("Signature", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	sigStr := ""
	sigStrBind := binding.BindString(&sigStr)
	sigEntry := widget.NewEntryWithData(sigStrBind)

	errStr := ""
	errStrBind := binding.BindString(&errStr)
	errEntry := widget.NewLabelWithData(errStrBind)

	signBtn := widget.NewButton("Sign", func() {
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
	})

	return container.NewVBox(
		head,
		privLb,
		privEntry,
		filePathLb,
		filePathEntry,
		sigLb,
		sigEntry,
		layout.NewSpacer(),
		errEntry,
		signBtn,
	), nil
}

func memzero(b []byte) {
	if b == nil {
		return
	}
	for i := range b {
		b[i] = 0
	}
}

func checkContainer() (*fyne.Container, error) {
	head := widget.NewLabelWithStyle("Signature verification", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	url := widget.NewLabelWithStyle("URL Ethereum node", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	urlEntry := widget.NewEntry()

	domain := widget.NewLabelWithStyle("Domain", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	domainEntry := widget.NewEntry()

	filePath := widget.NewLabelWithStyle("File path", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	filePathEntry := widget.NewEntry()

	signature := widget.NewLabelWithStyle("Signature", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	signatureEntry := widget.NewEntry()

	success := widget.NewLabelWithStyle("Success", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	success.Hide()

	invalid := widget.NewLabelWithStyle("Invalid", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	invalid.Hide()

	errStr := ""
	errStrBind := binding.BindString(&errStr)
	errEntry := widget.NewLabelWithData(errStrBind)

	btn := widget.NewButton("Sign", func() {
		errStrBind.Set("")
		errStrBind.Reload()

		success.Hide()
		invalid.Hide()

		client, err := ethclient.Dial(urlEntry.Text)
		if err != nil {
			errStrBind.Set("error: " + err.Error())
			errStrBind.Reload()
		}

		signatureByte, err := hex.DecodeString(signatureEntry.Text)
		if err != nil {
			errStrBind.Set("error: " + err.Error())
			errStrBind.Reload()
		}

		ok, err := check(client, domainEntry.Text, filePathEntry.Text, signatureByte)
		if err != nil {
			errStrBind.Set("error: " + err.Error())
			errStrBind.Reload()
		}
		if ok {
			success.Show()
		} else {
			invalid.Show()
		}
	})

	return container.NewVBox(
		head,
		url,
		urlEntry,
		domain,
		domainEntry,
		filePath,
		filePathEntry,
		signature,
		signatureEntry,
		layout.NewSpacer(),
		success,
		errEntry,
		btn,
	), nil
}

func check(client *ethclient.Client, domain, filepath string, signature []byte) (bool, error) {
	hash, err := filehash.FileHash(filepath)
	if err != nil {
		return false, err
	}
	address, err := signfile.HashPubKey(hash, signature)
	if err != nil {
		return false, err
	}
	check, err := ens.CheckEnsAddress(client, domain, address)
	if err != nil {
		return false, err
	}
	return check, nil
}
