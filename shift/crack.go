package shift

import (
	"bytes"
	"errors"
)

func Crack(ciphertext []byte, crib []byte) ([]byte, error) {
	for i := range 256 {
		v := Decipher([]byte{byte(i)}, ciphertext)
		if bytes.HasPrefix(v, crib) {
			return []byte{byte(i)}, nil
		}
	}
	return nil, errNotFound
}

var errNotFound = errors.New("not found")
