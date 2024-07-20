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
	crib := flag.String("crib", "", "Prefix to look for")
	outputKey := flag.Bool("output-key", false, "if set, will output key instead of the decoded value")
	detailed := flag.Bool("output-detailed", false, "if set, will output a more detailed summary")
	flag.Parse()
	if *crib == "" {
		fmt.Println("-crib is required")
		os.Exit(1)
	}
	ciphertext, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	key, err := shift.Crack(ciphertext[:len(*crib)], []byte(*crib))
	if err != nil {
		fmt.Printf("Failed to crack: %v\n", err)
		os.Exit(1)
	}
	plaintext := make([]byte, len(ciphertext))
	block, err := shift.NewShiftCipher(key)
	if err != nil {
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}
	if *outputKey {
		os.Stdout.Write([]byte(hex.EncodeToString(key)))
		os.Exit(0)
	}
	mode := shift.NewDecrypter(block)
	mode.CryptBlocks(plaintext, ciphertext)
	plaintext = shift.Unpad(plaintext)
	if *detailed {
		os.Stdout.WriteString("Key\t" + hex.EncodeToString(key))
		os.Stdout.WriteString("Plaintext:\n" + string(plaintext))
	} else {
		os.Stdout.Write(plaintext)
	}
}
