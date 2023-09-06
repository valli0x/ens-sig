package filehash

import (
	"io"
	"os"

	"golang.org/x/crypto/sha3"
)

func FileHash(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := sha3.NewLegacyKeccak256()

	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}
