package ics

import (
	"log"
	"net/http"
	"strings"
	"time"

	ical "github.com/arran4/golang-ical"
)

type Event struct {
	Summary     string
	Description string
	Start       string
	End         string
}

func ParseRemotely(u string) error {
	// u := "https://calendars.icloud.com/holiday/cn_zh.ics"
	resp, err := http.Get(u)
	if err != nil {
		return err
	}

	c, err := ical.ParseCalendar(resp.Body)
	if err != nil {
		return err
	}
	for _, component := range c.Components {
		ls := component.UnknownPropertiesIANAProperties()
		// b, _ := json.Marshal(ls)
		// log.Printf("%s", b)
		fs := map[string]string{}
		for _, v := range ls {
			fs[v.IANAToken] = v.Value
		}

		if strings.Contains(fs["SUMMARY"], "休") ||
			strings.Contains(fs["SUMMARY"], "假期") ||
			strings.Contains(fs["SUMMARY"], "班") {

			startDate := fs["DTSTART"]
			if len(startDate) > 8 {
				startDate = startDate[:8]
			}

			endDate := fs["DTEND"]
			if len(endDate) > 8 {
				endDate = endDate[:8]
			}

			layout := "20060102" // 时间字符串的格式

			// 将字符串转换为时间
			startTime, err := time.Parse(layout, startDate)
			if err != nil {
				log.Println(err)
				continue
			}

			endTime := startTime.Add(time.Second)
			if endDate != "" && endDate != startDate {
				endTime, err = time.Parse(layout, endDate)
				if err != nil {
					log.Println(err)
					continue
				}
			}
			// log.Println(endDate, endTime)

			// log.Println(">=", startDate, "<", endDate, ":", fs["SUMMARY"])
			for i := startTime; i.Before(endTime); i = i.AddDate(0, 0, 1) {
				log.Println(i.Format(layout), ">>", fs["SUMMARY"], "raw:>=", fs["DTSTART"], "<", fs["DTEND"])
			}

		}
	}
	return err
}
