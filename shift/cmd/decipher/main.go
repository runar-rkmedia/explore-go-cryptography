package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
)

func main() {
	key := flag.Int("key", 1, "Key to use for decipher")
	flag.Parse()
	ciphertext, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	plaintext := shift.Decipher(byte(*key), ciphertext)
	os.Stdout.Write(plaintext)
}
