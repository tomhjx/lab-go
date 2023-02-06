package list

import (
	"math/rand"
	"time"

	"log"
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

func Strings2Interfaces(l []string) []interface{} {
	ret := make([]interface{}, len(l))
	for i, v := range l {
		ret[i] = v
	}
	return ret
}

func RandomKey(l map[int64]int) int64 {
	weightSum := 0
	ll := [][]interface{}{}
	scopei0 := 0
	for k, v := range l {
		weightSum += v
		scopei2 := scopei0 + v
		ll = append(ll, []interface{}{k, scopei0, scopei2})
		scopei0 = scopei2
	}
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(weightSum)

	for _, v := range ll {
		if v[1].(int) <= randomNum && v[2].(int) > randomNum {
			return v[0].(int64)
		}
	}
	return 0
}

func RandomKeyn(l map[int64]int, count int) (ret []int64) {
	weightSum := 0
	ll := [][]interface{}{}
	scopei0 := 0
	for k, v := range l {
		weightSum += v
		scopei2 := scopei0 + v
		ll = append(ll, []interface{}{k, scopei0, scopei2})
		scopei0 = scopei2
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		randomNum := rand.Intn(weightSum)
		log.Printf("randomNum %v", randomNum)
		for _, v := range ll {
			if v[1].(int) <= randomNum && v[2].(int) > randomNum {
				ret = append(ret, v[0].(int64))
				break
			}
		}
	}

	return ret
}

func StickTopStrings(l []string, topl []string) []string {
	l1 := []string{}
	l2 := []string{}
	toplm := map[string]bool{}
	l1m := map[string]bool{}

	for _, v := range topl {
		toplm[v] = true
	}

	for _, v := range l {
		if _, ok := toplm[v]; ok {
			l1m[v] = true
			continue
		}
		l2 = append(l2, v)
	}
	for _, v := range topl {
		if _, ok := l1m[v]; ok {
			l1 = append(l1, v)
		}
	}

	return append(l1, l2...)
}
