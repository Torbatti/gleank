package utils

// ExistInSlice checks whether a comparable element exists in a slice of the same type.
func ExistInSlice[T comparable](item T, list []T) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}

	return false
}
