package shift

import (
	"crypto/cipher"
	"errors"
	"fmt"
)

const BlockSize = 32

type shiftCipher struct {
	key [BlockSize]byte
}

// Ensure struct implements cipher.Block (but don't return the interface, return the struct)
var _ cipher.Block = shiftCipher{}

// BlockSize implements cipher.Block.
func (c shiftCipher) BlockSize() int {
	return BlockSize
}

// Decrypt implements cipher.Block.
func (c shiftCipher) Decrypt(dst []byte, src []byte) {
	keyLength := len(src)
	for i := 0; i < len(src); i++ {
		dst[i] = src[i] - c.key[i%keyLength]
	}
	// panic("unimplemented")
}

func NewShiftCipher(key []byte) (*shiftCipher, error) {
	if len(key) != BlockSize {
		return nil, fmt.Errorf("%w: expected key of size %d, but received key of size %d", ErrKeySize, BlockSize, len(key))
	}
	return &shiftCipher{key: [BlockSize]byte(key)}, nil
}

var ErrKeySize = errors.New("invalid key size")

func (c shiftCipher) Encrypt(dst, src []byte) {
	keyLength := len(src)
	for i := 0; i < len(src); i++ {
		dst[i] = src[i] + c.key[i%keyLength]
	}
}

func Encipher(key []byte, plaintext []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	length := len(key)
	for i, b := range plaintext {
		ciphertext[i] = b + key[i%length]
	}
	return ciphertext
}

func Decipher(key []byte, ciphertext []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	length := len(key)
	if length == 0 {
		panic("NO KEY SPECIFIED")
	}
	for i, b := range ciphertext {
		plaintext[i] = b - key[i%length]
	}
	return plaintext
}
