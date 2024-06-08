package shift_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/runar-rkmedia/explore-go-cryptography/shift"
)

func TestShiftHappyPath(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		key         byte
		input, want []byte
	}{
		{
			key:   1,
			input: []byte("HAL"),
			want:  []byte("IBM"),
		},
		{
			key:   2,
			input: []byte("SPEC"),
			want:  []byte("URGE"),
		},
		{
			key:   3,
			input: []byte("PERK"),
			want:  []byte("SHUN"),
		},
		{
			key:   4,
			input: []byte("GEL"),
			want:  []byte("KIP"),
		},
		{
			key:   7,
			input: []byte("CHEER"),
			want:  []byte("JOLLY"),
		},
		{
			key:   10,
			input: []byte("BEEF"),
			want:  []byte("LOOP"),
		},
		{
			key:   1,
			input: []byte("ADD"),
			want:  []byte("BEE"),
		},
		{
			key:   1,
			input: []byte("ANA"),
			want:  []byte("BOB"),
		},
		{
			key:   1,
			input: []byte("INKS"),
			want:  []byte("JOLT"),
		},
		{
			key:   1,
			input: []byte("ADMIX"),
			want:  []byte("BENJY"),
		},
		{
			key:   1,
			input: []byte{0, 1, 2, 3, 255},
			want:  []byte{1, 2, 3, 4, 0},
		},
	}

	for _, tc := range tcs {
		name := fmt.Sprintf("%s + %d to %s", tc.input, tc.key, tc.want)
		t.Run(name, func(t *testing.T) {
			got := shift.Encipher(tc.key, tc.input)
			if !bytes.Equal(tc.want, got) {
				t.Errorf("shift.Encipher(%s) want %q, got %q", tc.input, tc.want, got)
			}
		})

	}

}
