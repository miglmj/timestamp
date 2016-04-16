package main

import (
	"encoding/json"
	"fmt"
	"github.com/miguelmejiamontes/timestamp/timeutils"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", timefunc)
	http.ListenAndServe(GetPort(), nil)
}

func timefunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		markdown, _ := ioutil.ReadFile("README.md")
		homepage := blackfriday.MarkdownBasic(markdown)
		w.Write(homepage)
		return
	}
	path := strings.TrimPrefix(r.URL.Path, "/")

	parsedString := timeutils.ParseTimestring(path)

	obj, _ := json.Marshal(parsedString)
	w.Write([]byte(obj))
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
