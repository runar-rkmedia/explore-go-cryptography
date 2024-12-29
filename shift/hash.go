package shift

import (
	"encoding/binary"
)

func LenHash(b []byte) []byte {
	out := make([]byte, 8)
	binary.BigEndian.PutUint64(out, uint64(len(b)))
	return out
}
