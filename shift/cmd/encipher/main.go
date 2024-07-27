package main

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
	"golang.org/x/term"
)

func main() {
	keyStr := flag.String("key", "01", "Key to use in hex form")
	outputb64url := flag.Bool("output-base-64-url", false, "if set, will output in base64 (urlencoded)")
	outputb64 := flag.Bool("output-base-64", false, "if set, will output in base64 (standard)")
	flag.Parse()
	key, err := hex.DecodeString(*keyStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	block, err := shift.NewShiftCipher(key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	plaintext, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	enc, iv, err := shift.NewCBCEncryptorEncrypterWithRandomIV(block)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	plaintext = shift.Pad(plaintext, enc.BlockSize())
	ciphertext := make([]byte, len(plaintext))
	enc.CryptBlocks(ciphertext, plaintext)
	if !*outputb64url && !*outputb64 {
		if term.IsTerminal(int(os.Stdout.Fd())) {
			os.Stdout.WriteString("terminal is interactive, forcing base64-output\n")
			*outputb64 = true
		}
	}
	if *outputb64url {
		out := append(iv, ciphertext...)
		os.Stdout.WriteString(base64.URLEncoding.EncodeToString(out))
	} else if *outputb64 {
		out := append(iv, ciphertext...)
		os.Stdout.WriteString(base64.StdEncoding.EncodeToString(out))
	} else {
		os.Stdout.Write(iv)
		os.Stdout.Write(ciphertext)
	}
}
