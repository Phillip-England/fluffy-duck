package utility

import "strconv"

func IsStringNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}