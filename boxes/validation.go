package boxes

import (
	"encoding/hex"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/valli0x/ens-sig/filehash"
	"github.com/valli0x/ens-sig/signfile"
	ens "github.com/wealdtech/go-ens/v3"
)

const (
	checkNameBox = "check"
)

func CheckContainer(w fyne.Window, back *widget.Button) (name string, _ *fyne.Container, _ error) {
	head := widget.NewLabelWithStyle("Signature verification", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	url := widget.NewLabelWithStyle("URL Ethereum node", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	urlEntry := widget.NewEntry()

	domain := widget.NewLabelWithStyle("Domain", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	domainEntry := widget.NewEntry()

	filePath := widget.NewLabelWithStyle("File", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	filePathBtn := ""
	filePathEntry := widget.NewButton("open file", func() {
		dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
			if uc != nil {
				filePathBtn = uc.URI().Path()
			}
		}, w)
	})

	signature := widget.NewLabelWithStyle("Signature", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	signatureEntry := widget.NewEntry()

	success := widget.NewLabelWithStyle("Success", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	success.Hide()

	invalid := widget.NewLabelWithStyle("Invalid", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	invalid.Hide()

	errStr := ""
	errStrBind := binding.BindString(&errStr)
	errEntry := widget.NewLabelWithData(errStrBind)

	btn := widget.NewButton("check", func() {
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

		ok, err := check(client, domainEntry.Text, filePathBtn, signatureByte)
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

	checkBox := container.NewVBox(
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
		back,
	)

	return checkNameBox, checkBox, nil
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
	resolved, err := ens.ReverseResolve(client, common.HexToAddress(address))
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(resolved, domain), nil
}
