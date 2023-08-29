package ens

import (
	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
)

// func example() {

// 	domain := "crypto-yoga.eth"
// 	address, err := ens.Resolve(client, domain)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Address of %s is %s\n", domain, address.Hex())

// 	reverse, err := ens.ReverseResolve(client, address)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if reverse == "" {
// 		fmt.Printf("%s has no reverse lookup\n", address.Hex())
// 	} else {
// 		fmt.Printf("Name of %s is %s\n", address.Hex(), reverse)
// 	}
// }

func CheckEnsAddress(client *ethclient.Client, domain, addressPub string) (bool, error) {
	address, err := ens.Resolve(client, domain)
	if err != nil {
		return false, err
	}
	return address.Hex() == addressPub, nil
}
