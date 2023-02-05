package list

import (
	"math/rand"
	"time"
)

// 数组分片
func ChunkInts(items []int, chunkSize int) (chunks [][]int) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

// 数组去重
func UniqueStrings(l []string) (ret []string) {
	exists := map[string]bool{}
	for _, v := range l {
		if _, ok := exists[v]; ok {
			continue
		}
		exists[v] = true
		ret = append(ret, v)
	}
	return ret
}

// 随机打乱数组
func Shuffle(l []interface{}) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(l), func(i, j int) {
		l[i], l[j] = l[j], l[i]
	})
}
