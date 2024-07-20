package shift

import (
	"crypto/cipher"
)

type decrypter struct {
	block     cipher.Block
	blockSize int
}

// Ensure struct implements cipher.BlockMode (but don't return the interface, return the struct)
var _ cipher.BlockMode = decrypter{}

// BlockSize implements cipher.BlockMode.
func (e decrypter) BlockSize() int {
	return BlockSize
}

// CryptBlocks implements cipher.BlockMode.
func (e decrypter) CryptBlocks(dst []byte, src []byte) {
	if len(src)%e.blockSize != 0 {
		panic("decrypter: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("decrypter: output smaller than input")
	}
	for len(src) > 0 {
		e.block.Decrypt(dst[:e.blockSize], src[:e.blockSize])
		src = src[e.blockSize:]
		dst = dst[e.blockSize:]
	}
}

func NewDecrypter(block cipher.Block) *decrypter {
	return &decrypter{block: block, blockSize: block.BlockSize()}
}
