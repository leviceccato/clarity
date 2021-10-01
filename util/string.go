package util

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
	return string(Subrune([]rune(str), start, length))
}

func Subrune(runes []rune, start, length int) []rune {
	if start >= len(runes) {
		return []rune{}
	}
	if start+length > len(runes) {
		length = len(runes) - start
	}
	return runes[start : start+length]
}
