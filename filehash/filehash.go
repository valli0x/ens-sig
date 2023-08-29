package filehash

import (
	"io"
	"os"

	"golang.org/x/crypto/sha3"
)

// FileHash return an ethereum format hash from a file
func FileHash(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := sha3.NewLegacyKeccak256() // this func uses to hash ethereum tx in the go-ethereum package

	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}
