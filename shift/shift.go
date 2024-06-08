package shift

func Encipher(plaintext []byte) []byte {
	b := make([]byte, len(plaintext))

	// return nil
	for i, v := range plaintext {
		b[i] = v + 1
	}
	return b
}
