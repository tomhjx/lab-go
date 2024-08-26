package random

import (
	"math/rand"
	"testing"
	"time"
)

func TestGetIntn(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				max: 100,
			},
		},
	}

	rand.Seed(time.Now().UnixNano())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := map[int]int{}
			for i := 0; i < 100; i++ {
				// got := GetIntn(tt.args.max)
				got := rand.Intn(tt.args.max)
				cs[got]++
			}
			t.Logf("len:%v", len(cs))
			for k, v := range cs {
				if v > 1 {
					t.Logf("%v count = %v", k, v)
				}
			}

			// for i := 0; i < 100; i++ {
			// 	t.Logf("rand:%v", rand.Intn(tt.args.max))
			// }
		})
	}
}
