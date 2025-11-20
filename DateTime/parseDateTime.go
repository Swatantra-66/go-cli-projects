package main

import (
	"fmt"
	"log"
	"time"
)

func parseDateTime() {
	timeString := "2025-11-15 12:03:08"
	layout := "2006-01-02 03:04:05"
	parsedTime, err := time.Parse(layout, timeString)
	if err != nil {
		log.Fatalf("Unable to parse time due to %v.", err)
	}

	fmt.Println("Parsed time is", parsedTime)

}
