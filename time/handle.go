package time

import "time"

func HoursUnix(n time.Duration) int64 {
	return time.Now().Add(n * time.Hour).Truncate(time.Hour).Unix()
}

func DaysUnix(n int) int64 {
	return DaysTime(n).Unix()
}

func DaysYmd(n int) string {
	return DaysTime(n).Format("20060102")
}

func MonthsYmd(n int) string {
	return MonthsTime(n).Format("20060102")
}
func DaysTime(n int) time.Time {
	return time.Now().AddDate(0, 0, n)
}

func MonthsTime(n int) time.Time {
	return time.Now().AddDate(0, n, 0)
}

func FirstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func LastDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).AddDate(0, 1, -1)
}
