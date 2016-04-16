package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type timestamp struct {
	Unix  string `json:"unix"`
	Human string `json:"natural"`
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")

	parsedString := parseTimestring(path)

	obj, _ := json.Marshal(parsedString)
	w.Write([]byte(obj))
}

func parseTimestring(timestring string) timestamp {
	humanForm := "January 01, 2016"
	if _, err := strconv.Atoi(timestring); err != nil {
		t, _ := time.Parse(humanForm, timestring)
		unixtime := t.Format(time.UnixDate)
		return timestamp{timestring, unixtime}
	} else {
		t, _ := time.Parse(time.UnixDate, timestring)
		humantime := t.Format(humanForm)
		return timestamp{humantime, timestring}
	}
}
