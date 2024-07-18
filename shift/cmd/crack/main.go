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
	if *detailed {
		os.Stdout.WriteString("Key\t" + hex.EncodeToString(key))
		plaintext := shift.Decipher(key, ciphertext)
		os.Stdout.WriteString("Plaintext:\n" + string(plaintext))
	} else if *outputKey {
		os.Stdout.Write([]byte(hex.EncodeToString(key)))
	} else {
		plaintext := shift.Decipher(key, ciphertext)
		os.Stdout.Write(plaintext)

	}
}
