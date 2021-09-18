package utility

// Find the unique strings in the first slice between 2 slices
// of strings
func SliceStringDifference(a, b []string) []string {
	uniqueStrings := []string{}
	isStringUnique := true
	for _, stringA := range a {
		for _, stringB := range b {
			if stringA == stringB {
				isStringUnique = false
			}
		}
		if isStringUnique {
			uniqueStrings = append(uniqueStrings, stringA)
		}
		isStringUnique = true
	}
	return uniqueStrings
}

// Have to convert to runes to perform a substr so that
// it supports non-ASCII characters
func Substr(str string, start, length int) string {
	runes := []rune(str)
	if start >= len(runes) {
		return ""
	}
	if start+length > len(runes) {
		length = len(runes) - start
	}
	return string(runes[start : start+length])
}
