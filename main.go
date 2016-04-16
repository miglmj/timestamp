package main

import (
	"encoding/json"
	"github.com/miguelmejiamontes/timestamp/timeutils"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", timefunc)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func timefunc(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")

	parsedString := timeutils.ParseTimestring(path)

	obj, _ := json.Marshal(parsedString)
	w.Write([]byte(obj))
}
