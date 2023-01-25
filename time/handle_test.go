package time

import (
	"testing"
)

func TestHoursUnix(t *testing.T) {
	t.Logf("after 1 hour unix: %v", HoursUnix(1))
}

func TestDaysUnix(t *testing.T) {
	t.Logf("after 1 day unix: %v", DaysUnix(1))
}

func TestDaysYmd(t *testing.T) {
	t.Logf("after 1 day Ymd: %v", DaysYmd(1))
}
