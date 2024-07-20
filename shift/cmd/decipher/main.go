package main

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
)

func main() {
	keyStr := flag.String("key", "01", "Key to use in hex form")
	inputBase64 := flag.Bool("input-base-64", false, "use if the input is base64-encoded")
	flag.Parse()
	key, err := hex.DecodeString(*keyStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	c, err := shift.NewShiftCipher(key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)

	}
	ciphertext, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *inputBase64 {
		l, err := base64.StdEncoding.Decode(ciphertext, ciphertext)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		ciphertext = ciphertext[:l]
	}
	plaintext := make([]byte, len(ciphertext))
	c.Decrypt(plaintext, ciphertext)
	os.Stdout.Write(plaintext)
}
