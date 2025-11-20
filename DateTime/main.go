package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Printf("Current(local) time is %v\n", currentTime)
	fmt.Printf("UTC time is %v\n", currentTime.UTC().Format("2006-01-02 03:04:05")) //UTC time
	fmt.Println(currentTime.Format("2006-1-2 15:4:5 pm"))                           //formatted date and time
	fmt.Println(currentTime.Format(time.RFC3339Nano))                               //the time.RFC3339Nano version includes nanoseconds in the format.
	fmt.Println("Formatted Time:", currentTime.Format("2006/01/02 03:04 PM"))
	fmt.Printf("The year is %v\n", currentTime.Year())
	fmt.Printf("The month is %v\n", currentTime.Month())
	fmt.Printf("The weekday is %v\n", currentTime.Weekday())
	fmt.Printf("The day is %v\n", currentTime.Day())
	fmt.Printf("The hour is %v\n", currentTime.Hour())
	fmt.Printf("The minute is %v\n", currentTime.Hour())
	fmt.Printf("The second is %v\n", currentTime.Second())

	//time arithmetic
	futureTime := currentTime.Add(10 * 24 * time.Hour) //added 10 days to the current time
	fmt.Printf("Future time: %v\n", futureTime.Format("2006-01-02 03:04 PM"))

	pastTime := currentTime.Add(-6 * 24 * time.Hour) //subtracted 6 days to the current time
	fmt.Printf("Past time: %v\n", pastTime.Format("2006-01-02 03:04:05 PM"))

	updatedDate := currentTime.AddDate(4, 2, 10) // Add 4 year, 2 months, and 10 days
	fmt.Println("Updated date: ", updatedDate.Format("2006-01-02"))

	//convert between time zones
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Error loading location ", err)
		return
	}
	India := currentTime.In(loc)
	fmt.Println("Time in India:", India.Format("2006-01-02 03:04:05"))

	parseDateTime() //task:parse date and time

}
