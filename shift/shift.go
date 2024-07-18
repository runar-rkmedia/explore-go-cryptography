package shift

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
