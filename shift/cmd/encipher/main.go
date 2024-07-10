package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
)

func main() {
	key := flag.Int("key", 1, "Key to use for encypher")
	flag.Parse()
	plaintext, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	ciphertext := shift.Encipher([]byte{byte(*key)}, plaintext)
	os.Stdout.Write(ciphertext)
}
