package shift

import (
	"crypto/cipher"
	"crypto/rand"
)

type encrypter struct {
	block     cipher.Block
	blockSize int
}

// Ensure struct implements cipher.BlockMode (but don't return the interface, return the struct)
var _ cipher.BlockMode = encrypter{}

// BlockSize implements cipher.BlockMode.
func (e encrypter) BlockSize() int {
	return BlockSize
}

// CryptBlocks implements cipher.BlockMode.
func (e encrypter) CryptBlocks(dst []byte, src []byte) {
	if len(src)%e.blockSize != 0 {
		panic("encrypter: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("encrypter: output smaller than input")
	}
	for len(src) > 0 {
		e.block.Encrypt(dst[:e.blockSize], src[:e.blockSize])
		src = src[e.blockSize:]
		dst = dst[e.blockSize:]
	}
}

// Deprecated
func NewEncrypter(block cipher.Block) cipher.BlockMode {
	return &encrypter{block: block, blockSize: block.BlockSize()}
}

func NewCBCEncryptorEncrypterWithRandomIV(block cipher.Block) (cipher.BlockMode, []byte, error) {
	iv := make([]byte, block.BlockSize())
	_, err := rand.Read(iv)
	if err != nil {
		return nil, nil, err
	}
	return cipher.NewCBCEncrypter(block, iv), iv, nil
}

func NewShiftCBCEncryptorEncrypterWithRandomIV(key []byte) (cipher.BlockMode, []byte, error) {
	block, err := NewShiftCipher(key)
	if err != nil {
		return nil, nil, err
	}
	return NewCBCEncryptorEncrypterWithRandomIV(block)
}
