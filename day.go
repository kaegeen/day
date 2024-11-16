package main

import (
	"fmt"
	"time"
)

func main() {
	// Specify the year you want to count days for
	year := 2024

	// Create a map to store the counts for each day of the week
	dayCounts := make(map[string]int)

	// Loop through all the days of the year
	for month := 1; month <= 12; month++ {
		// Get the number of days in the month
		daysInMonth := getDaysInMonth(year, month)
		for day := 1; day <= daysInMonth; day++ {
			// Create a date object for each day
			date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			// Get the weekday for the current date
			weekday := date.Weekday().String()
			// Increment the count for that weekday
			dayCounts[weekday]++
		}
	}

	// Print out the count for each day of the week
	for _, day := range []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"} {
		fmt.Printf("%s: %d\n", day, dayCounts[day])
	}
}

// Helper function to get the number of days in a month
func getDaysInMonth(year int, month int) int {
	// Create a date for the first day of the next month
	date := time.Date(year, time.Month(month)+1, 1, 0, 0, 0, 0, time.UTC)
	// Subtract one day to get the last day of the month
	lastDayOfMonth := date.Add(-24 * time.Hour)
	// Return the day of the month (which is the number of days in the month)
	return lastDayOfMonth.Day()
}
