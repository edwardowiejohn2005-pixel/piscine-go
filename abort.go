package piscine

func Abort(a, b, c, d, e int) int {
	numbers := []int{a, b, c, d, e}

	// Simple bubble sort (manual sorting)
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	// The median is the middle value
	return numbers[2]
}
