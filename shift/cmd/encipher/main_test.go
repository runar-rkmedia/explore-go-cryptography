package main

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
	"testing"
)

func runEncrypt(reader io.Reader, args ...string) ([]byte, error) {
	argus := append([]string{"run", "."}, args...)
	cmd := exec.Command("go", argus...)
	cmd.Stdin = reader
	out, err := cmd.CombinedOutput()
	return out, err
}

func runDecrypt(reader io.Reader, args ...string) ([]byte, error) {
	argus := append([]string{"run", "../decipher/"}, args...)
	cmd := exec.Command("go", argus...)
	cmd.Stdin = reader
	out, err := cmd.CombinedOutput()
	return out, err
}

var testkey = "0101010101010101010101010101010101010101010101010101010101010101"

func TestCliEncipherRequireKey(t *testing.T) {
	t.Parallel()
	var err error
	out, err := runEncrypt(nil)
	if err == nil {
		t.Fatalf("Expected error, but got success-code with output: %s", out)
	}
	if !strings.Contains(string(out), "invalid key size") {
		t.Fatalf("Expected output to include information indicating invalid key size but got: %s", out)
	}
}

func TestCliEncipherReportShortKey(t *testing.T) {
	t.Parallel()
	var err error
	out, err := runEncrypt(nil, "-key", "12345678")
	if err == nil {
		t.Fatalf("Expected error, but got success-code with output: %s", out)
	}
	if !strings.Contains(string(out), "invalid key size") {
		t.Fatalf("Expected output to include information indicating invalid key size but got: %s", out)
	}
	if !strings.Contains(string(out), "key of size 4") {
		t.Fatalf("Expected output to include information indicating invalid key size but got: %s", out)
	}
}

func TestCliEncipherWorksForSimpleCases(t *testing.T) {
	t.Parallel()
	plainTexts := [][]byte{
		[]byte("foobar"),
		[]byte(""),
		bytes.Repeat([]byte("foobar"), 13),
	}
	for _, plainText := range plainTexts {
		r := bytes.NewReader(plainText)
		encryptedOutput, err := runEncrypt(r, "-key", testkey)
		if err != nil {
			t.Fatal(encryptedOutput)
		}
		toDecrypt := bytes.NewReader(encryptedOutput)
		decryptedOutput, err := runDecrypt(toDecrypt, "-key", testkey)
		if err != nil {
			t.Fatal(encryptedOutput)
		}

		if !bytes.Equal(decryptedOutput, []byte(plainText)) {
			t.Fatalf("\nwant\t%v\n got\t%v", plainText, string(decryptedOutput))
		}
	}
}

func TestCliEncipherWorksForSimpleCase(t *testing.T) {
	plainText := []byte("foobar")
	r := bytes.NewReader(plainText)
	encryptedOutput, err := runEncrypt(r, "-key", testkey)
	if err != nil {
		t.Fatal(encryptedOutput)
	}
	wantCiphertext := []byte("gppcbs")
	if !bytes.HasPrefix(encryptedOutput, wantCiphertext) {
		t.Fatalf("\nwant\t%v\n got\t%v", wantCiphertext, encryptedOutput)
	}
	toDecrypt := bytes.NewReader(encryptedOutput)
	decryptedOutput, err := runDecrypt(toDecrypt, "-key", testkey)
	if err != nil {
		t.Fatal(encryptedOutput)
	}

	if !bytes.Equal(decryptedOutput, []byte(plainText)) {
		t.Fatalf("\nwant\t%v\n got\t%v", plainText, string(decryptedOutput))
	}
}
