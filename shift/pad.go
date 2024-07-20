package shift

import "bytes"

func Pad(input []byte, blockSize int) []byte {
	padding := blockSize - (len(input) % blockSize)
	return append(input, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func Unpad(input []byte) []byte {
	padding := int(input[len(input)-1])
	return input[:len(input)-padding]
}
