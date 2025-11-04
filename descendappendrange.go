package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{} // return empty slice, not nil
	}

	result := []int{} // initialize as empty slice literal

	for i := max; i > min; i-- {
		result = append(result, i)
	}

	return result
}
