package shift

import (
	"crypto/cipher"
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

func NewEncrypter(block cipher.Block) *encrypter {
	return &encrypter{block: block, blockSize: block.BlockSize()}
}
