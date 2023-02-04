package list

import (
	"testing"
)

func TestChunkInt(t *testing.T) {
	l := []int{}
	for i := 0; i < 55; i++ {
		l = append(l, i)
	}
	lg := ChunkInts(l, 10)
	for _, v := range lg {
		t.Logf("%v", v)
	}
}
