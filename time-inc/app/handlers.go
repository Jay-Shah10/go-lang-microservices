package app

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	// We are going to check if 'tz' is there. If so, then we are going to split it by ',' to make an array.
	// From here - for each location that is placed we print out the item.
	if r.URL.Query().Has("tz") {
		tzs := r.URL.Query().Get("tz")
		tzs_split := strings.Split(tzs, ",")

		for _, tz := range tzs_split {
			location, _ := time.LoadLocation(tz)
			fmt.Printf("Current Time: %s: %s\n", time.Now().In(location), location)
		}
		//location, _ := time.LoadLocation(tz)
		//fmt.Printf("Current Time: %s: %s\n", time.Now().In(location), location)

	} else {
		fmt.Printf("current_time: %s\n", time.Now().UTC())
	}
	// TODO: convert the output to JSON format.
	// TODO: creat a struct that we can be used to make it into JSON format.

}
