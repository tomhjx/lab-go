package context

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type index struct {
	Version   string
	NameSpace string
	Group     string
	Key       string
}

func TestStore(t *testing.T) {
	ctx := context.Background()
	k1 := "k1"
	v1 := "v1"
	ctx = SetValue(ctx, "k1", "v1")
	r := getFromBgCtx(k1)
	t.Logf("getFromBgCtx: %v", r)
	assert.Equal(t, nil, r)

	r = getFromCtx(ctx, k1)
	t.Logf("getFromCtx: %v", r)
	assert.Equal(t, v1, r)

	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		r := getFromBgCtx(k1)
		t.Logf("getFromBgCtx in goroutine(1): %v", r)
		go func() {
			defer wg.Done()
			r := getFromBgCtx(k1)
			t.Logf("getFromBgCtx in goroutine(1/1): %v", r)
		}()
	}()
	go func() {
		defer wg.Done()
		r = getFromCtx(ctx, k1)
		t.Logf("getFromCtx in goroutine(2): %v", r)
	}()

	go func() {
		defer wg.Done()
		r = getFromCtx(context.Background(), k1)
		t.Logf("getFromCtx in goroutine(3): %v", r)
	}()

	wg.Wait()

	index1 := index{Version: "v1", NameSpace: "default", Group: "user", Key: "id"}
	ctx = SetValue(ctx, index1, "v1Ofindex1")
	r = getFromCtx(ctx, index1)
	t.Logf("getFromCtx with index1: %v", r)

	index2 := index{Version: "v1", NameSpace: "default", Group: "user", Key: "id"}
	ctx = SetValue(ctx, index2, "v1Ofindex2")
	r = getFromCtx(ctx, index2)
	t.Logf("getFromCtx with index2: %v", r)

}

func getFromBgCtx(key any) any {
	return GetValue(context.Background(), key)
}

func getFromCtx(ctx context.Context, key any) any {
	return GetValue(ctx, key)
}
