package main

import (
	"fmt"
)

// Weather () - displays the weather based on the queryResult
func Weather(res NLPResponse) {
	fmt.Println(res.Response)
}
