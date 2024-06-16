package shift

import (
	"bytes"
	"errors"
)

func Crack(ciphertext []byte, crib []byte) (byte, error) {
	for i := range 256 {
		v := Decipher(byte(i), ciphertext)
		if bytes.HasPrefix(v, crib) {
			return byte(i), nil
		}
	}
	return 0, errNotFound
}

var errNotFound = errors.New("not found")
