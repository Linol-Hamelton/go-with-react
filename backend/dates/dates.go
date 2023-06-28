package dates

import (
	"time"
)

func GetDates() (string, string) {
	// Get the current date and time
	now := time.Now()

	// Get the first day of the previous month
	firstDay := time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, time.UTC)

	// Get the last day of the previous month
	lastDay := firstDay.AddDate(0, 1, -1)

	// Format the dates in the specified format
	firstDayFormatted := firstDay.Format("2006-01-02")
	lastDayFormatted := lastDay.Format("2006-01-02")

	return firstDayFormatted, lastDayFormatted
}

/*
func main() {
	firstDay, lastDay := GetDates()
	fmt.Println(firstDay, lastDay)
}
*/
