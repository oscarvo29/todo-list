package main

import (
	"errors"
	"fmt"
	"time"
)

/*
	TODO:
	- implement an opptional column, for which there is a deadline.
	- Consider to make the table more pretty with column borders, row borders, and so on.
*/

func main() {
	// db.InitDB()
	// args := os.Args

	// if len(args) <= 1 {
	// 	return
	// }

	// services.ProcessCmd(args...)

	timeStr := []string{
		"01-01",
		"01:14",
		"24-06-2025 01:00",
		"28-02-2025 09:10",
		"29/03/1998 11:32",
		"tomorrow",
	}

	for idx, str := range timeStr {
		t, err := GetTimeFromString(str)

		fmt.Println("")
		fmt.Printf("Input: %s, Index: %d \n", str, idx)
		fmt.Printf("Time: %v, Err: %s", t, err)
		fmt.Println("")
	}
}

func GetTimeFromString(input string) (time.Time, error) {
	layouts := []string{
		"01-01-2024",
		"01-01",
		"01/01/2024",
		"01/01",
		"02-01-2006 15:04",
		"01-01 15:04",
		"02/01/2006 15:04",
		"01/01 15:04",
		"15:04",
	}

	if input == "tomorrow" {
		tomorrow := time.Now().AddDate(0, 0, 1)
		return tomorrow, nil
	}

	for _, layout := range layouts {
		parsedTime, err := time.Parse(layout, input)
		if err == nil {
			result := GetFormat(layout, parsedTime)
			return result, nil
		}
	}

	return time.Now(), errors.New("not the right date time formation")
}

func GetFormat(layout string, parsedTime time.Time) time.Time {
	now := time.Now()
	switch layout {
	case "01-01", "01/01", "01-01 15:04", "01/01 15:04":
		parsedDate := time.Date(now.Year(), parsedTime.Month(), parsedTime.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, now.Location())
		if parsedDate.After(now) {
			return parsedDate
		}

		parsedDate = parsedDate.AddDate(1, 0, 0)
		return parsedDate
	case "15:04":
		result := time.Date(now.Year(), now.Month(), now.Day()+1,
			parsedTime.Hour(), parsedTime.Minute(), 0, 0, now.Location())
		return result
	default:
		return parsedTime
	}
}
