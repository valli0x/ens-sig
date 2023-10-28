package signfile

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func SignHash(hash []byte, privateKey string) ([]byte, error) {
	privECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	sig, err := crypto.Sign(hash, privECDSA)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func HashPubKey(hash, sig []byte) (address string, _ error) {
	pubSigBytes, err := crypto.Ecrecover(hash, sig)
	if err != nil {
		return "", err
	}
	pubSigECDSA, err := crypto.UnmarshalPubkey(pubSigBytes)
	if err != nil {
		return "", err
	}
	return crypto.PubkeyToAddress(*pubSigECDSA).Hex(), nil
}
