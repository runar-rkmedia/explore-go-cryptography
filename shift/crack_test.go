package shift_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
)

func TestCrack(t *testing.T) {
	t.Parallel()
	for _, tc := range tcs {
		name := fmt.Sprintf("%s + %d = %s", tc.plaintext, tc.key,
			tc.ciphertext)
		t.Run(name, func(t *testing.T) {
			got, err := shift.Crack(tc.ciphertext, tc.plaintext[:3])
			if err != nil {
				t.Fatalf("Failed to crack %s: %v", tc.ciphertext, err)
			}
			if !bytes.Equal(tc.key, got) {
				t.Fatalf("want %d, got %d", tc.key, got)
			}
		})
	}
}
