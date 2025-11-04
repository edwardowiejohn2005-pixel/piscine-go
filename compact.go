package piscine

func Compact(ptr *[]string) int {
	slice := *ptr // dereference the pointer to access the actual slice
	newSlice := []string{}

	for _, str := range slice {
		if str != "" {
			newSlice = append(newSlice, str)
		}
	}

	*ptr = newSlice // update the original slice through the pointer
	return len(newSlice)
}
