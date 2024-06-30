package piscine

func isSpace(c rune) bool {
    return c == ' ' || c == '\f' || c == '\n' || c == '\r' || c == '\t' || c == '\v'
}

func Atoi(s string) int {
	result := 0
	sign := 1
	started := false

	for _, char := range s {
		if !started && isSpace(char) {
			continue
		}
		if !started && (char == '-' || char == '+') {
			if char == '-' {
				sign = -1
			}
			started = true
			continue
		}
		
		if char < '0' || char > '9' {
			return 0
		}

		digit := int(char - '0')
		result = result*10 + digit
		started = true
	}
	return result * sign
}
