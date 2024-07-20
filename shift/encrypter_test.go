package shift_test

import (
	"bytes"
	"testing"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
)

func Test_EncrypterEncipheresBlockAlignedMessage(t *testing.T) {
	t.Parallel()
	plaintext := []byte("This message is exactly 32 bytes")
	block, err := shift.NewShiftCipher(testKey)
	if err != nil {
		t.Fatal(err)
	}
	enc := shift.NewEncrypter(block)
	want := []byte("Uijt!nfttbhf!jt!fybdumz!43!czuft")
	got := make([]byte, 32)
	enc.CryptBlocks(got, plaintext)
	if !bytes.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func Test_EncrypterEncipheresBlockMisalignedMessage(t *testing.T) {
	t.Parallel()
	plaintext := []byte("This message is exactly 32 bytes, and then some")
	block, err := shift.NewShiftCipher(testKey)
	if err != nil {
		t.Fatal(err)
	}
	enc := shift.NewEncrypter(block)
	want := []byte("Uijt!nfttbhf!jt!fybdumz!43!czuft")
	got := make([]byte, 32)
	enc.CryptBlocks(got, plaintext)
	if !bytes.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func Test_EncrypterCorrectlyReportsBlockSize(t *testing.T) {
	t.Parallel()
	block, err := shift.NewShiftCipher(testKey)
	if err != nil {
		t.Fatal(err)
	}
	enc := shift.NewEncrypter(block)
	want := shift.BlockSize
	got := enc.BlockSize()
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
