package time

import "time"

func NextHourlyUnix() int64 {
	return time.Now().Add(1 * time.Hour).Truncate(time.Hour).Unix()
}
