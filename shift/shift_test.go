package shift_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
)

var tcs = []struct {
	key                   byte
	plaintext, ciphertext []byte
}{
	{
		key:        1,
		plaintext:  []byte("HAL"),
		ciphertext: []byte("IBM"),
	},
	{
		key:        2,
		plaintext:  []byte("SPEC"),
		ciphertext: []byte("URGE"),
	},
	{
		key:        3,
		plaintext:  []byte("PERK"),
		ciphertext: []byte("SHUN"),
	},
	{
		key:        4,
		plaintext:  []byte("GEL"),
		ciphertext: []byte("KIP"),
	},
	{
		key:        7,
		plaintext:  []byte("CHEER"),
		ciphertext: []byte("JOLLY"),
	},
	{
		key:        10,
		plaintext:  []byte("BEEF"),
		ciphertext: []byte("LOOP"),
	},
	{
		key:        1,
		plaintext:  []byte("ADD"),
		ciphertext: []byte("BEE"),
	},
	{
		key:        1,
		plaintext:  []byte("ANA"),
		ciphertext: []byte("BOB"),
	},
	{
		key:        1,
		plaintext:  []byte("INKS"),
		ciphertext: []byte("JOLT"),
	},
	{
		key:        1,
		plaintext:  []byte("ADMIX"),
		ciphertext: []byte("BENJY"),
	},
	{
		key:        1,
		plaintext:  []byte{0, 1, 2, 3, 255},
		ciphertext: []byte{1, 2, 3, 4, 0},
	},
}

func TestEncipher(t *testing.T) {
	t.Parallel()

	for _, tc := range tcs {
		name := fmt.Sprintf("%s + %d to %s", tc.plaintext, tc.key, tc.ciphertext)
		t.Run(name, func(t *testing.T) {
			got := shift.Encipher(tc.key, tc.plaintext)
			if !bytes.Equal(tc.ciphertext, got) {
				t.Errorf("shift.Encipher(%s) want %q, got %q", tc.plaintext, tc.ciphertext, got)
			}
		})

	}
}

func TestDecipher(t *testing.T) {
	t.Parallel()
	for _, tc := range tcs {
		name := fmt.Sprintf("%s - %d = %s", tc.plaintext, tc.key, tc.ciphertext)
		t.Run(name, func(t *testing.T) {
			got := shift.Decipher(tc.key, tc.ciphertext)
			if !bytes.Equal(tc.plaintext, got) {
				t.Errorf("shift.Decipher(%s) want %q, got %q", tc.plaintext, tc.ciphertext, got)
			}
		})

	}
}
