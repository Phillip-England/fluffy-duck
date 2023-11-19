package utility

import "strconv"

func StringToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}