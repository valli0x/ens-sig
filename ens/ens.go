package ens

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
)

func CheckEnsAddress(client *ethclient.Client, domain, addressPub string) (bool, error) {
	resolved, err := ens.ReverseResolve(client, common.HexToAddress(addressPub))
	fmt.Println(resolved)
	if err != nil {
		return false, err
	}
	return resolved == domain, nil
}
