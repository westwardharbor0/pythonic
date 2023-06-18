package pythonic

// IN checks if values is present in given slice.
func IN[T comparable](slice []T, value T) bool {
	return IndexOf(slice, value) != -1
}

// IndexOf returns index of the given value in slice.
// If multiple values are present the first index will be returned.
func IndexOf[T comparable](slice []T, value T) int {
	for index, item := range slice {
		if item == value {
			return index
		}
	}
	return -1
}

// Set returns only unique items in slice.
func Set[T comparable](slice []T) []T {
	result := make([]T, 0, len(slice))
	for _, item := range slice {
		if !IN(result, item) {
			result = append(result, item)
		}
	}
	return result
}

// Remove deletes item by value from slice.
// If value not found input slice is returned.
// If value has multiple occurrences in slice only first is deleted.
func Remove[T comparable](slice []T, value T) []T {
	if indexOf := IndexOf(slice, value); indexOf >= 0 {
		return append(slice[:indexOf], slice[indexOf+1:]...)
	}
	return slice
}

// Counter returns amount of occurrences of each item in string slice.
func Counter(slice []string) map[string]int {
	result := make(map[string]int, 0)
	for _, item := range slice {
		result[item] += 1
	}
	return result
}
