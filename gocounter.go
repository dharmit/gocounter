package main

import (
	"expvar"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/paulbellamy/ratecounter"
)

var (
	counter       *ratecounter.RateCounter
	hitsperminute = expvar.NewInt("hits per minute")
	html_resp     = "<div align=center>" + "<img src='https://pbs.twimg.com/profile_images/" +
		"458352291767013376/K9nN_rhH_400x400.png'>" +
		"<h1>This page has been visited %s times in paste one minute!</h1>" +
		"<br>" +
		"</div>"
)

func increment(w http.ResponseWriter, r *http.Request) {
	counter.Incr(1)
	hitsperminute.Set(counter.Rate())
	fmt.Fprintf(w, html_resp, strconv.FormatInt(counter.Rate(), 10))
}

func main() {
	counter = ratecounter.NewRateCounter(1 * time.Minute)
	http.HandleFunc("/", increment)
	http.ListenAndServe(":80", nil)
}
