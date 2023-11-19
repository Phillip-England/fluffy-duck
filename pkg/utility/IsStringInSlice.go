package utility

func IsStringInSlice(allowed []string, value string) bool {
	for _, v := range allowed {
		if value == v {
			return true
		}
	}
	return false
}