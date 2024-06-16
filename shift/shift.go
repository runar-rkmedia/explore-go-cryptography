package shift

func Encipher(key byte, plaintext []byte) []byte {
	b := make([]byte, len(plaintext))

	// return nil
	for i, v := range plaintext {
		b[i] = v + key
	}
	return b
}

func Decipher(key byte, ciphertext []byte) []byte {
	return Encipher(-key, ciphertext)
}
