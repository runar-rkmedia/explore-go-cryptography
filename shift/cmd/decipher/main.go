package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
)

func main() {
	keyStr := flag.String("key", "01", "Key to use in hex form")
	flag.Parse()
	key, err := hex.DecodeString(*keyStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	ciphertext, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	plaintext := shift.Decipher(key, ciphertext)
	os.Stdout.Write(plaintext)
	// fmt.Println("ciphertext", string(ciphertext), string(plaintext), len(plaintext))
}
