package time

import (
	"testing"
	"time"
)

func TestHoursUnix(t *testing.T) {
	t.Logf("after 1 hour unix: %v", HoursUnix(1))
}

func TestDaysUnix(t *testing.T) {
	t.Logf("after 1 day unix: %v", DaysUnix(1))
}

func TestDaysYmd(t *testing.T) {
	t.Logf("after 1 day Ymd: %v", DaysYmd(1))
	t.Logf("after -2 day Ymd: %v", DaysYmd(-2))
	t.Logf("after -4 day Ymd: %v", DaysYmd(-4))
}

func TestMonthsYmd(t *testing.T) {
	t.Logf("after 1 month Ymd: %v", MonthsYmd(1))
	t.Logf("after -2 month Ymd: %v", MonthsYmd(-2))
	t.Logf("after -4 month Ymd: %v", MonthsYmd(-4))
}

func TestLastDayOfMonth(t *testing.T) {
	ps := []string{
		"20230111",
		"20230212",
		"20231113",
	}
	for _, v := range ps {
		te, _ := time.Parse("20060102", v)
		t.Logf("last day of %s: %v", v[:6], LastDayOfMonth(te).Format("2006-01-02"))
	}
}
func TestFirstDayOfMonth(t *testing.T) {
	ps := []string{
		"20230111",
		"20230212",
		"20231113",
	}
	for _, v := range ps {
		te, _ := time.Parse("20060102", v)
		t.Logf("first day of %s: %v", v[:6], FirstDayOfMonth(te).Format("2006-01-02"))
	}
}

func TestAddDate(t *testing.T) {
	now := time.Now()
	ls := []struct {
		txt    string
		fn     func(time.Time) time.Time
		layout string
	}{
		{
			txt: "first day of the mouth",
			fn: func(te time.Time) time.Time {
				return time.Date(te.Year(), te.Month(), 1, 0, 0, 0, 0, te.Location())
			},
			layout: "2006-01-02",
		},
		{
			txt: "last day of the mouth",
			fn: func(te time.Time) time.Time {
				return time.Date(te.Year(), te.Month(), 0, 0, 0, 0, 0, te.Location()).AddDate(0, 1, 0)
			},
			layout: "2006-01-02",
		},
		{
			txt: "last day of the last mouth",
			fn: func(te time.Time) time.Time {
				return time.Date(te.Year(), te.Month(), 0, 0, 0, 0, 0, te.Location())
			},
			layout: "2006-01-02",
		},
	}
	for _, v := range ls {
		t.Logf("%s : %s", v.txt, v.fn(now).Format(v.layout))
	}

}
func TestFormat(t *testing.T) {
	t.Logf("now Y-m-d H: %s", time.Now().Format("2006-01-02 15"))
}
