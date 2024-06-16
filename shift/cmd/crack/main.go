package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
)

func main() {
	crib := flag.String("crib", "", "Prefix to look for")
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
		fmt.Println("Failed to crack: %w", err)
		os.Exit(1)
	}
	plaintext := shift.Decipher(key, ciphertext)
	os.Stdout.Write(plaintext)
}
