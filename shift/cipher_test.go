package shift_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
)

var testKey = bytes.Repeat([]byte{1}, shift.BlockSize)

var cipherCases = []struct {
	plaintext, ciphertext []byte
}{
	{
		plaintext:  []byte{0, 1, 2, 3, 4, 5},
		ciphertext: []byte{1, 2, 3, 4, 5, 6},
	},
}

var tcs = []struct {
	key                   []byte
	plaintext, ciphertext []byte
}{
	{
		key:        []byte{1},
		plaintext:  []byte("HAL"),
		ciphertext: []byte("IBM"),
	},
	{
		key:        []byte{2},
		plaintext:  []byte("SPEC"),
		ciphertext: []byte("URGE"),
	},
	{
		key:        []byte{3},
		plaintext:  []byte("PERK"),
		ciphertext: []byte("SHUN"),
	},
	{
		key:        []byte{4},
		plaintext:  []byte("GEL"),
		ciphertext: []byte("KIP"),
	},
	{
		key:        []byte{7},
		plaintext:  []byte("CHEER"),
		ciphertext: []byte("JOLLY"),
	},
	{
		key:        []byte{10},
		plaintext:  []byte("BEEF"),
		ciphertext: []byte("LOOP"),
	},
	{
		key:        []byte{1},
		plaintext:  []byte("ADD"),
		ciphertext: []byte("BEE"),
	},
	{
		key:        []byte{1},
		plaintext:  []byte("ANA"),
		ciphertext: []byte("BOB"),
	},
	{
		key:        []byte{1},
		plaintext:  []byte("INKS"),
		ciphertext: []byte("JOLT"),
	},
	{
		key:        []byte{1},
		plaintext:  []byte("ADMIX"),
		ciphertext: []byte("BENJY"),
	},
	{
		key:        []byte{1},
		plaintext:  []byte{0, 1, 2, 3, 255},
		ciphertext: []byte{1, 2, 3, 4, 0},
	},
	{
		key:        []byte{1, 2, 3},
		plaintext:  []byte{0, 0, 0},
		ciphertext: []byte{1, 2, 3},
	},
	{
		key:        []byte{1, 2},
		plaintext:  []byte{0, 1, 2},
		ciphertext: []byte{1, 3, 3},
	},
}

func TestNewCipher_GivesErrorForEmptyKey(t *testing.T) {
	t.Parallel()
	_, err := shift.NewShiftCipher([]byte{})
	if !errors.Is(err, shift.ErrKeySize) {
		t.Fatalf("a zero-length key should not be accepted, got an unexpected error: %v", err)
	}
}

func TestNewCipher_GivesNoErrorForValidKey(t *testing.T) {
	t.Parallel()
	_, err := shift.NewShiftCipher(make([]byte, shift.BlockSize))
	if err != nil {
		t.Fatalf("got an unexpected error: %v", err)
	}
}

func TestNewCipher_GivesErrorForInvalidKeyLength(t *testing.T) {
	t.Parallel()
	_, err := shift.NewShiftCipher(make([]byte, 7))
	if !errors.Is(err, shift.ErrKeySize) {
		t.Fatalf("a zero-length key should not be accepted, got an unexpected error: %v", err)
	}
}

func Test_shiftCipher_Encrypt(t *testing.T) {
	t.Parallel()
	for _, tc := range cipherCases {
		name := fmt.Sprintf("%s -> %s", tc.plaintext, tc.ciphertext)
		c, err := shift.NewShiftCipher(testKey)
		if err != nil {
			t.Fatalf("want no error in call to shift.NewShiftCipher, got %v", err)
		}
		t.Run(name, func(t *testing.T) {
			got := make([]byte, len(tc.plaintext))
			c.Encrypt(got, tc.plaintext)
			if !bytes.Equal(tc.ciphertext, got) {
				t.Errorf("Encrypt(%s) want %q, got %q", tc.plaintext, tc.ciphertext, got)
				for i := 0; i < len(tc.ciphertext); i++ {
					if tc.ciphertext[i] != got[i] {
						t.Logf("First Incorrect byte at index %d:  %q != %q", i, tc.ciphertext[i], got[i])
						break
					}
				}
			}
		})
	}
}

func Test_shiftCipher_Decrypt(t *testing.T) {
	t.Parallel()
	for _, tc := range cipherCases {
		name := fmt.Sprintf("%s -> %s", tc.ciphertext, tc.plaintext)
		c, err := shift.NewShiftCipher(testKey)
		if err != nil {
			t.Fatalf("want no error in call to shift.NewShiftCipher, got %v", err)
		}
		t.Run(name, func(t *testing.T) {
			got := make([]byte, len(tc.ciphertext))
			c.Decrypt(got, tc.ciphertext)
			if !bytes.Equal(tc.plaintext, got) {
				t.Errorf("Decrypt(%s) want %q, got %q", tc.plaintext, tc.plaintext, got)
				for i := 0; i < len(tc.plaintext); i++ {
					if tc.plaintext[i] != got[i] {
						t.Logf("First Incorrect byte at index %d:  %q != %q", i, tc.plaintext[i], got[i])
						break
					}
				}
			}
		})
	}
}
