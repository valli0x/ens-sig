package boxes

import (
	"encoding/hex"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
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

	filePath := widget.NewLabelWithStyle("File path", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	filePathEntry := widget.NewEntry()

	signature := widget.NewLabelWithStyle("Signature", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	signatureEntry := widget.NewEntry()

	success := widget.NewLabelWithStyle("Valid", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	success.Hide()

	invalid := widget.NewLabelWithStyle("Invalid", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	invalid.Hide()

	resolve := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	resolve.Hide()

	errStr := ""
	errStrBind := binding.BindString(&errStr)
	errEntry := widget.NewLabelWithData(errStrBind)

	btn := widget.NewButton("check", func() {
		errStrBind.Set("")
		errStrBind.Reload()

		success.Hide()
		invalid.Hide()
		resolve.Hide()

		client, err := ethclient.Dial(urlEntry.Text)
		if err != nil {
			errStrBind.Set("eth connect error: " + err.Error())
			errStrBind.Reload()
			return
		}

		signatureByte, err := hex.DecodeString(signatureEntry.Text)
		if err != nil {
			errStrBind.Set("decode error: " + err.Error())
			errStrBind.Reload()
			return
		}

		resolved, ok, err := check(client, domainEntry.Text, filePathEntry.Text, signatureByte)
		if err != nil {
			errStrBind.Set("check signature error: " + err.Error())
			errStrBind.Reload()
			invalid.Show()
			return
		}

		if ok {
			success.Show()
		} else {
			invalid.Show()
			resolve.SetText(resolved)
			resolve.Show()
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
		invalid,
		resolve,
		errEntry,
		btn,
		back,
	)

	return checkNameBox, checkBox, nil
}

func check(client *ethclient.Client, domain, filepath string, signature []byte) (string, bool, error) {
	hash, err := filehash.FileHash(filepath)
	if err != nil {
		return "", false, err
	}
	address, err := signfile.HashPubKey(hash, signature)
	if err != nil {
		return "", false, err
	}
	resolved, err := ens.ReverseResolve(client, common.HexToAddress(address))
	if err != nil {
		return "", false, err
	}
	return resolved, reflect.DeepEqual(resolved, domain), nil
}
