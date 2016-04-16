package main

import (
	"encoding/json"
	"fmt"
	"github.com/miguelmejiamontes/timestamp/timeutils"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", timefunc)
	http.ListenAndServe(GetPort(), nil)
}

func timefunc(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")

	parsedString := timeutils.ParseTimestring(path)

	obj, _ := json.Marshal(parsedString)
	w.Write([]byte(obj))
}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
