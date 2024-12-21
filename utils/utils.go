package utils

import (
	"fmt"
	"strconv"
	"time"
)

func ConvertStringToInt64(inp string) (int64, error) {
	return strconv.ParseInt(inp, 10, 64)
}

func CompareTimes(t time.Time) {
	nowTime := time.Now()

	difference := nowTime.Sub(t)

	if difference < time.Minute {
		fmt.Println("seconds ago")
	} else if difference < time.Hour {
		fmt.Println("created less than an hour ago..")
	} else if difference < 24*time.Hour {
		fmt.Println("created less than a day ago")
	} else {
		// days := int(difference.Hours() / 24)
		fmt.Println("Happen some days ago")
	}
}
