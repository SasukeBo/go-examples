package algorithm

// UniqueAppend append item and remove duplicated one
func UniqueAppend(slice []int, item int) []int {
	for i, j := range slice {
		if j == item {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}

	slice = append(slice, item)
	return slice
}
