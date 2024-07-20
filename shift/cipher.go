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

var (
	// Ensure struct implements cipher.Block (but don't return the interface, return the struct)
	_          cipher.Block = shiftCipher{}
	ErrKeySize error        = errors.New("invalid key size")
)

func NewShiftCipher(key []byte) (*shiftCipher, error) {
	if len(key) != BlockSize {
		return nil, fmt.Errorf("%w: expected key of size %d, but received key of size %d", ErrKeySize, BlockSize, len(key))
	}
	return &shiftCipher{key: [BlockSize]byte(key)}, nil
}

func (c shiftCipher) BlockSize() int {
	return BlockSize
}

func (c shiftCipher) Decrypt(dst []byte, src []byte) {
	keyLength := len(c.key)
	for i := 0; i < len(src); i++ {
		dst[i] = src[i] - c.key[i%keyLength]
	}
	// panic("unimplemented")
}

func (c shiftCipher) Encrypt(dst, src []byte) {
	keyLength := len(c.key)
	for i := 0; i < len(src); i++ {
		dst[i] = src[i] + c.key[i%keyLength]
	}
}
