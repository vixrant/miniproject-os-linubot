package main

import (
	"fmt"
	"time"
)

// Weather - displays the weather based on the queryResult
func Weather(res NLPResponse) error {
	fmt.Println(res.Response)

	// First we get the time from the response
	var target time.Time
	if res.Entities["date"] != "" {
		var err error
		target, err = time.Parse(time.RFC3339, res.Entities["date"])
		if err != nil {
			return err
		}
	} else {
		target = time.Now()
	}
	fmt.Println(target)

	// TODO: Make a POST request to Openweather API

	return nil
}
