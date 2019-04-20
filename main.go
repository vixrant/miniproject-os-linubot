package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	query := strings.Join(args[1:], " ")
	if query == "" {
		fmt.Println("No query passed")
		return
	}

	DF.init("v-dj-loc-19", "keys/v-dj-loc-19-d70c3550e084.json", "en", "Asia/Colombo")

	// Use NLP
	response := DF.processNLP(query, "not-important-session-id")
	// fmt.Printf("%#v", response)
	analyseResponse(response)
}

func analyseResponse(res NLPResponse) {
	switch intent := res.Intent; intent {
	case "WeatherIntent":
		err := Weather(res)
		if err != nil {
			fmt.Println(err)
		}
	case "WelcomeIntent":
		fmt.Println("Hi! Try weather, filesearch, commandsearch. ")

	case "FileSearchIntent":
		ImageSearchByDate(res)

	case "CommandIntent":
		CommandSearch(res)

	case "GrepIntent":
		GrepSearch(res)

	case "TldrIntent":
		TldrSearch(res)

	default:
		fmt.Println("SORRY! Couldn't recognise this action!")
	}
}
