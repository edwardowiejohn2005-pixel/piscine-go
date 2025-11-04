package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}

	var result []int
	for _, ch := range str {
		result = append(result, int(ch))
	}
	return result
}
