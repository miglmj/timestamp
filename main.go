package main

import (
	"encoding/json"
	"github.com/miglmj/timestamp/timeutils"
	"github.com/miglmj/tools"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", timefunc)
	http.ListenAndServe(tools.GetPort(), nil)
}

func timefunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write(tools.RenderHome())
		return
	}
	path := strings.TrimPrefix(r.URL.Path, "/")

	parsedString := timeutils.ParseTimestring(path)

	obj, _ := json.Marshal(parsedString)
	w.Write([]byte(obj))
}
