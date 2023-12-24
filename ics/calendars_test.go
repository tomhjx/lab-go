package ics

import "testing"

func TestParseRemotely(t *testing.T) {

	// u := "https://www.shuyz.com/githubfiles/china-holiday-calender/master/holidayCal.ics"
	u := "https://calendars.icloud.com/holiday/cn_zh.ics"
	t.Log(ParseRemotely(u))
}
