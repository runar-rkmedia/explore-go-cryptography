package shift_test

import (
	"bytes"
	"testing"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
)

func Test_DecrypterDecryptsBlockAlignedMessage(t *testing.T) {
	t.Parallel()
	ciphertext := []byte("Uijt!nfttbhf!jt!fybdumz!43!czuft")
	block, err := shift.NewShiftCipher(testKey)
	if err != nil {
		t.Fatal(err)
	}
	enc := shift.NewDecrypter(block)
	want := []byte("This message is exactly 32 bytes")
	got := make([]byte, 32)
	enc.CryptBlocks(got, ciphertext)
	if !bytes.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func Test_DecrypterCorrectlyReportsBlockSize(t *testing.T) {
	t.Parallel()
	block, err := shift.NewShiftCipher(testKey)
	if err != nil {
		t.Fatal(err)
	}
	enc := shift.NewDecrypter(block)
	want := shift.BlockSize
	got := enc.BlockSize()
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
