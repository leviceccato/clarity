package util

// Enforce panic for functions where err should be nil (programmer error)
func MustGet[T any](value T, err error) T {
	Must(err)
	return value
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Find unique items between two slices
func Unique[T comparable](aItems, bItems []T) []T {
	items := make([]T, 0, len(aItems))
	isUnique := true

	for _, a := range aItems {
		for _, b := range bItems {
			if a == b {
				isUnique = false
			}
		}

		if isUnique {
			items = append(items, a)
		}

		isUnique = true
	}

	return items
}
