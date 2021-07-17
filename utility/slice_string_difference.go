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
