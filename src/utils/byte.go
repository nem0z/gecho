package utils

import (
	"crypto/rand"
	"crypto/sha256"
)

func Checksum(data []byte) []byte {
	hash := sha256.Sum256(data)
	hash = sha256.Sum256(hash[:])
	return hash[:4]
}

func Nonce(len int) ([]byte, error) {
	nonce := make([]byte, len)
	_, err := rand.Read(nonce)

	return nonce, err
}
