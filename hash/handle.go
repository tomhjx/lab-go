package hash

import (
	"crypto/sha256"
	"encoding/binary"
)

func GetUInt64(s string, b int) uint64 {
	h := sha256.Sum256([]byte(s))
	return binary.BigEndian.Uint64(h[:b])
}
