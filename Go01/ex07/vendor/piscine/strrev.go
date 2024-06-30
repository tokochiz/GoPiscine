package piscine

func strLen(s string) int {
	l := 0
	for range s {
		l++
	}
	return l
}

func StrRev(s string) string {
	l := strLen(s)
	rev := make([]rune, l)
	for i, char := range s {
		rev[l-1-i] = char
	}

	return string(rev)
}
