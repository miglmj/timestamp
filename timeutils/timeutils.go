package timeutils

import (
	"strconv"
	"time"
)

type dualtime struct {
	Unix  string `json:"unix"`
	Human string `json:"natural"`
}

func ParseTimestring(timestring string) dualtime {
	humanForm := "January 2, 2006"
	var unixtime string
	var humantime string

	if _, err := strconv.Atoi(string(timestring[0])); err != nil {
		t, _ := time.Parse(humanForm, timestring)
		unixtime = strconv.FormatInt(t.Unix(), 10)
		humantime = timestring
	} else {
		u, _ := strconv.ParseInt(timestring, 10, 64)
		t := time.Unix(u, 0)
		humantime = t.Format(humanForm)
	}

	return dualtime{unixtime, humantime}
}
