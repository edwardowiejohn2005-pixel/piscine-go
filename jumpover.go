package piscine

func JumpOver(str string) string {
	// Handle empty or too-short strings
	if str == "" || len(str) < 3 {
		return "\n"
	}

	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	return result + "\n"
}
