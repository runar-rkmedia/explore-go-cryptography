package shift_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
)

func TestCrack(t *testing.T) {
	t.Parallel()
	plaintext := []byte("This message is exactly 32 bytes")
	ciphertext := []byte("Uijs message is exactly 32 bytes")
	want := append([]byte{1, 1, 1}, bytes.Repeat([]byte{0}, 29)...)
	got, err := shift.Crack(ciphertext, plaintext)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func BenchmarkNext(b *testing.B) {
	key := bytes.Repeat([]byte{0}, 32)
	b.ResetTimer()
	for range b.N {
		key, _ = shift.Next(key)
	}
}

func TestNextCorrectlyIncrementsInputWithoutOverflow(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		input, want []byte
	}{
		{
			input: []byte{0, 0, 0},
			want:  []byte{1, 0, 0},
		},
		{
			input: []byte{255, 0, 0},
			want:  []byte{0, 1, 0},
		},
		{
			input: []byte{255, 255, 0},
			want:  []byte{0, 0, 1},
		},
		{
			input: []byte{255, 255, 254},
			want:  []byte{0, 0, 255},
		},
	}

	for _, tt := range tcs {
		t.Run(fmt.Sprintf("%x", tt.input), func(t *testing.T) {
			got, err := shift.Next(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(tt.want, got) {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestNextErrsWhenOverflowing(t *testing.T) {
	t.Parallel()
	_, err := shift.Next([]byte{255, 255, 255})
	if err == nil {
		t.Fatal("Expected error for overflow, but got nilw")
	}
}
