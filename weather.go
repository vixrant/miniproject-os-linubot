package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{}

// Weather - displays the weather based on the queryResult
func Weather(res NLPResponse) error {
	fmt.Println(res.Response)

	// First we get the time from the response
	// var target time.Time
	if res.Entities["date"] != "" {
		var err error
		_, err = time.Parse(time.RFC3339, res.Entities["date"].(string))
		if err != nil {
			return err
		}
		// TODO: Do something with date
	}

	req, err := http.NewRequest("GET", "http://wttr.in/mumbai", nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "curl")

	apiRes, err := client.Do(req)
	if err != nil {
		return err
	}
	defer apiRes.Body.Close()

	body, err := ioutil.ReadAll(apiRes.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}
