package utils

import "strconv"

func ConvertStringToInt64(inp string) (int64, error) {
	return strconv.ParseInt(inp, 10, 64)
}
