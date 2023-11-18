package utility

func ConditionalString(condition bool, someString string) string {
	if condition {
		return someString
	}
	return ""
}