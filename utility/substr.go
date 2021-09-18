package utility

// Convert string to runes
func Substr(str string, start int, length int) string {
	runes := []rune(str)
	if start >= len(runes) {
		return ""
	}
	if start+length > len(runes) {
		length = len(runes) - start
	}
	return string(runes[start : start+length])
}
