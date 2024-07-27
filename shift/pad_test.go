package shift_test

import (
	"bytes"
	"testing"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
)

var padCases = []struct {
	name        string
	raw, padded []byte
}{
	{
		name:   "1 short of full block",
		raw:    []byte{0, 0, 0},
		padded: []byte{0, 0, 0, 1},
	},
	{
		name:   "2 short of full block",
		raw:    []byte{0, 0},
		padded: []byte{0, 0, 2, 2},
	},
	{
		name:   "3 short of full block",
		raw:    []byte{0},
		padded: []byte{0, 3, 3, 3},
	},
	{
		name:   "Full block",
		raw:    []byte{0, 0, 0, 0},
		padded: []byte{0, 0, 0, 0, 4, 4, 4, 4},
	},
	{
		name:   "Empty block",
		raw:    []byte{},
		padded: []byte{4, 4, 4, 4},
	},
}

func Test_Pad(t *testing.T) {
	t.Parallel()
	for _, tt := range padCases {
		got := shift.Pad(tt.raw, 4)
		if !bytes.Equal(tt.padded, got) {
			t.Fatalf("want %d, got %d", tt.padded, got)
		}
	}
}

func Test_Unap(t *testing.T) {
	for _, tt := range padCases {
		got := shift.Unpad(tt.padded, 4)
		if !bytes.Equal(tt.raw, got) {
			t.Fatalf("want %d, got %d", tt.raw, got)
		}
	}
}
