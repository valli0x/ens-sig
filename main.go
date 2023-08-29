package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/valli0x/ens-sig/ens"
	"github.com/valli0x/ens-sig/filehash"
	"github.com/valli0x/ens-sig/signfile"
)

func main() {
	domain := "hello.eth"
	filepath := "example.txt"
	priv := "8981cbc60cbbfc9324091f7fe6826e63c853a91150a72b97b7e050404c958037"

	signature, err := sign(filepath, priv)
	if err != nil {
		panic(err)
	}

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/22f2cce9d2334104bb27152a63dbc4b5")
	if err != nil {
		panic(err)
	}

	signed, err := check(client, domain, filepath, signature)
	if err != nil {
		panic(err)
	}

	fmt.Println("signature check:", signed)
}

func sign(filepath string, privKey string) ([]byte, error) {
	// get hash file
	h, err := filehash.FileHash(filepath)
	if err != nil {
		return nil, err
	}

	// get signature
	signature, err := signfile.SignHash(h, privKey)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

func check(client *ethclient.Client, domain, filepath string, signature []byte) (bool, error) {
	// get hash file
	hash, err := filehash.FileHash(filepath)
	if err != nil {
		return false, err
	}

	address, err := signfile.HashPub(hash, signature)
	if err != nil {
		return false, err
	}

	check, err := ens.CheckEnsAddress(client, domain, address)
	if err != nil {
		return false, err
	}

	return check, nil
}
