package shift_test

import (
	"bytes"
	"testing"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
)

func TestLenHashReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	input := []byte("I love you, Bob")
	want := []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x0f,
	}
	got := shift.LenHash(input)
	t.Log(got)
	if !bytes.Equal(want, got) {
		t.Errorf("%s: want %x, got %x", input, want, got)
	}
}
