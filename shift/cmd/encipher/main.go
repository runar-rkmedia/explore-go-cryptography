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
	outputb64url := flag.Bool("output-base-64-url", false, "if set, will output in base64 (urlencoded)")
	outputb64 := flag.Bool("output-base-64", false, "if set, will output in base64 (standard)")
	flag.Parse()
	key, err := hex.DecodeString(*keyStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	plaintext, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	ciphertext := shift.Encipher(key, plaintext)
	if *outputb64url {
		os.Stdout.WriteString(base64.URLEncoding.EncodeToString(ciphertext))
	} else if *outputb64 {
		os.Stdout.WriteString(base64.StdEncoding.EncodeToString(ciphertext))
	} else {
		os.Stdout.Write(ciphertext)
	}
}
