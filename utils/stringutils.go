package utils

import "strconv"

func StrToInt(inValue string) int {
	val, err := strconv.Atoi(inValue)
	if err != nil {
		return 0
	}
	return val
}
