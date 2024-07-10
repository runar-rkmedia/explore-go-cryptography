package shift

func Encipher(key []byte, plaintext []byte) []byte {
	b := make([]byte, len(plaintext))

	// return nil
	l := len(key)
	for i, v := range plaintext {
		b[i] = v + key[i%l]
	}
	return b
}

func Decipher(key []byte, ciphertext []byte) []byte {
	return nil
	// return Encipher(-key, ciphertext)
}
