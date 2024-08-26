package hash

import (
	"testing"
	"time"
)

func TestGetUInt64(t *testing.T) {

	type args struct {
		s string
		b int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "test1",
			args: args{
				s: "aaa",
				b: 8,
			},
			want: 10967539898259627185,
		},
	}

	t.Logf("hour:%v\n", time.Now().Hour())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUInt64(tt.args.s, tt.args.b); got != tt.want {
				t.Errorf("GetUInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
