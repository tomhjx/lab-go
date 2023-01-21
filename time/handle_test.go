package time

import (
	"testing"
)

func TestNextHourlyUnix(t *testing.T) {
	t.Logf("time.Truncate: %v", NextHourlyUnix())
}
