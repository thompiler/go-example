package main

import (
	"regexp"

	"github.com/airbrake/gobrake/v4"
)

var blocklist = []interface{}{regexp.MustCompile("thingo"), regexp.MustCompile("bingo")}

var airbrake = gobrake.NewNotifierWithOptions(&gobrake.NotifierOptions{
	ProjectId:     123,            // Replace with project ID
	ProjectKey:    "abcdef123456", // Replace with project API key
	KeysBlocklist: blocklist,      // Change to "KeysBlocklist" to test deprecated option
})

func main() {
	defer airbrake.Close()

	notice := gobrake.NewNotice("operation failed", nil, 0)
	notice.Context["thingo"] = "filter me"
	notice.Context["bingo"] = "me too"
	notice.Context["ringo"] = "don't filter me out"
	airbrake.Notify(notice, nil)
}
