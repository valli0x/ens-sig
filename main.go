package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/valli0x/ens-sig/containers"
)

// domain := "hello.eth"
// filepath := "examples/f1.txt"
// priv := 8981cbc60cbbfc9324091f7fe6826e63c853a91150a72b97b7e050404c958037
// c3774ad1ba8885227c51e7ad5b40924aa4402d09ae725880455ffd0f615cff81378cb4137a30564f9c28632bf000fbc8587eec4fb63c45bddcc377646579bc6c01
// ethclientURL := "https://mainnet.infura.io/v3/22f2cce9d2334104bb27152a63dbc4b5"

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
