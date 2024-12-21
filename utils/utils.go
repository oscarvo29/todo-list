package utils

import (
	"fmt"
	"strconv"
	"time"
)

func ConvertStringToInt64(inp string) (int64, error) {
	return strconv.ParseInt(inp, 10, 64)
}

func CompareTimes(t time.Time) string {
	nowTime := time.Now()

	difference := nowTime.Sub(t)

	if difference < time.Minute {
		return "seconds ago"
	} else if difference < time.Hour {
		return "created less than an hour ago.."
	} else if difference < 24*time.Hour {
		return "created less than a day ago"
	} else {
		// days := int(difference.Hours() / 24)
		return "Happen some days ago"
	}
}

// PrintBlankSpaces makes sure each field in table, have the right spacing for so that each field in a row.
// This makes sure that each field of a row, is straight under each column header.
// This function takes in a integer and prints the needed blank spaces. If the needed blank spaces is 0, it adds 3.
// /his is to make sure there at least are some space between the longest word in a colunmn and the next column start.
func PrintBlankSpaces(blankSpaces int) {
	if blankSpaces == 0 {
		for range 3 {
			fmt.Print(" ")
		}
		return
	}

	blankSpaces = blankSpaces + 3
	for range blankSpaces {
		fmt.Print(" ")
	}
}
