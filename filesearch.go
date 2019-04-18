package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var allowedExts = [5]string{".png", ".raw", ".jpg", ".jpeg", ".gif"}

func contains(a [5]string, x string) bool {
	for _, e := range allowedExts {
		if e == x {
			return true
		}
	}
	return false
}

func analyse(path string, info os.FileInfo, startDate time.Time, endDate time.Time) error {
	if ext := filepath.Ext(path); !contains(allowedExts, ext) {
		return nil
	}
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}

	if date := stat.ModTime(); date.After(startDate) && date.Before(endDate) {
		fmt.Println("File: ", info.Name(), " - ", path, " - ModDate - ", stat.ModTime())
	}

	return nil
}

type datePeriod struct {
	startDate time.Time `json:"startDate"`
	endDate   time.Time `json:"endDate"`
}

// ImageSearchByDate - search files based on date modified.
func ImageSearchByDate(res NLPResponse) error {
	var dp datePeriod
	err := json.Unmarshal([]byte(res.Entities["date-period"]), &dp)
	if err != nil {
		fmt.Println(res.Entities["date-period"])
		return err
	}

	fmt.Println(dp)

	home, err := os.UserHomeDir()

	err = filepath.Walk(home,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			analyse(path, info, dp.startDate, dp.endDate)

			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return nil
}
