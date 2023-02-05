package list

import (
	"testing"
)

func TestChunkInt(t *testing.T) {
	l := []int{}
	for i := 0; i < 55; i++ {
		l = append(l, i)
	}

	sizes := []int{10, 100}
	for _, s := range sizes {
		t.Logf("chunk by size:%v", s)
		lg := ChunkInts(l, s)
		for _, v := range lg {
			t.Logf("%v", v)
		}
	}

}

func TestUniqueStrings(t *testing.T) {
	l := []string{"a", "b", "a", "c", "b", "d"}
	l = UniqueStrings(l)
	t.Logf("%v", l)
}

func TestShuffle(t *testing.T) {
	l := []interface{}{"a1", "a2", "a3", "b1", "b2", "b3"}
	Shuffle(l)
	t.Logf("%v", l)
}
