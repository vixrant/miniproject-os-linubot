package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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

// ImageSearchByDate - search files based on date modified.
func ImageSearchByDate(res NLPResponse) error {
	dp := res.Entities["date-period"].(string)
	dates := strings.Split(dp, " ")[1:3]
	startDate, err := time.Parse(time.RFC3339, dates[0])
	if err != nil {
		return err
	}
	endDate, err := time.Parse(time.RFC3339, dates[1])
	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()

	err = filepath.Walk(home,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			analyse(path, info, startDate, endDate)

			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return nil
}
