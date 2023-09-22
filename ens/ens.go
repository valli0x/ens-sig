package ens

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
)

func CheckEnsAddress(client *ethclient.Client, domain, addressPub string) (bool, error) {
	resolved, err := ens.ReverseResolve(client, common.HexToAddress(addressPub))
	if err != nil {
		return false, err
	}
	return resolved == domain, nil
}
