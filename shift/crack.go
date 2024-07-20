package shift

import (
	"bytes"
	"errors"
	"fmt"
)

const (
	MaxKeyLen = 32
)

func Crack(ciphertext []byte, crib []byte) (key []byte, err error) {
	plaintext := make([]byte, len(crib))
	key = make([]byte, BlockSize)
	// i := 0
	for {
		// i++
		// if i%1000 == 0 {
		// 	fmt.Print("\r", hex.EncodeToString(key))
		// }
		block, err := NewShiftCipher(key)
		if err != nil {
			panic(err)
		}
		block.Decrypt(plaintext, ciphertext[:len(crib)])
		if bytes.Equal(crib, plaintext) {
			return key, nil
		}
		key, err = Next(key)
		if err != nil {
			return nil, fmt.Errorf("No key found %w", err)
		}
	}
	// for k := range min(MaxKeyLen, len(ciphertext), len(crib)) {
	// 	for guess := range 256 {
	// 		result := ciphertext[k] - byte(guess)
	// 		if result == crib[k] {
	// 			key = append(key, byte(guess))
	// 			break
	// 		}
	// 	}
	// 	if len(key) > 0 {
	// 		padded := make([]byte, BlockSize)
	// 		copy(padded, key)
	// 		// for i := 0; i < min(BlockSize, len(key)); i++ {
	// 		// 	padded[BlockSize-i-1] = key[i]
	// 		// }
	// 		// fmt.Println("")
	// 		// fmt.Println(padded)
	// 		// fmt.Println(key)
	// 		block, err := NewShiftCipher(padded)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		dec := NewDecrypter(block)
	//
	// 		// block.Decrypt(padded, ciphertext[:len(crib)])
	// 		dec.CryptBlocks(padded, ciphertext[:len(crib)])
	// 		if bytes.Equal(crib, padded) {
	// 			return key, nil
	// 		}
	// 	}
	// 	// if len(key) > 0 && bytes.Equal(crib, Decipher(key, ciphertext[:len(crib)])) {
	// 	// 	return key, nil
	// 	// }
	// }
	// return nil, errNotFound
}

func Next(key []byte) ([]byte, error) {
	for i := range key {
		if key[i] < 255 {
			key[i]++
			return key, nil
		}
		key[i] = 0
	}
	return nil, errors.New("overflow")
}

var errNotFound = errors.New("not found")
