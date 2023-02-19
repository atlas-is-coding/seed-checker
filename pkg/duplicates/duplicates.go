package duplicates

func RemoveFromSlice[T comparable](arr []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0, len(arr))

	for _, val := range arr {
		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}

	return result
}
