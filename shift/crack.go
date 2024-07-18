package shift

import (
	"bytes"
	"errors"
)

const (
	MaxKeyLen = 32
)

func Crack(ciphertext []byte, crib []byte) (key []byte, err error) {
	for k := range min(MaxKeyLen, len(ciphertext), len(crib)) {
		for guess := range 256 {
			result := ciphertext[k] - byte(guess)
			if result == crib[k] {
				key = append(key, byte(guess))
				break
			}
		}
		if len(key) > 0 && bytes.Equal(crib, Decipher(key, ciphertext[:len(crib)])) {
			return key, nil
		}
	}
	return nil, errNotFound
}

var errNotFound = errors.New("not found")
