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

func TestShuffleAny(t *testing.T) {
	l := []interface{}{"a1", "a2", "a3", "b1", "b2", "b3"}
	Shuffle(l)
	t.Logf("%v", l)
}

func TestShuffleStrings(t *testing.T) {
	l := []string{"a1", "a2", "a3", "b1", "b2", "b3"}
	l1 := Strings2Interfaces(l)
	Shuffle(l1)
	t.Logf("%v", l)
	t.Logf("%v", l1)
}

func TestRandomKey(t *testing.T) {
	w := map[int64]int{}
	w[1] = 1
	w[2] = 1
	// w[3] = 1

	c := map[int64]int{}

	for i := 0; i < 100; i++ {
		c[RandomKey(w)]++
	}
	t.Logf("%v", c)
	t.Logf("once random key %v", RandomKey(w))
}

func TestRandomKeyn(t *testing.T) {
	w := map[int64]int{}
	w[1] = 1
	w[2] = 1
	// w[3] = 2

	ll := RandomKeyn(w, 100)
	c := map[int64]int{}
	for _, v := range ll {
		c[v]++
	}
	t.Logf("%v", c)
}

func TestStickTopStrings(t *testing.T) {
	topl := []string{"US", "CN", "JP", "IN", "ID"}
	l := []string{"AR", "BR", "CL", "CO", "FI", "ID", "IL", "MX", "PH", "PR", "PT", "VN", "US"}
	l = StickTopStrings(l, topl)
	t.Logf("%v", l)
}
