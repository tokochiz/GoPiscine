package piscine

import "ft"

func PrintStr(s string) {
	for _, char := range s {
		ft.PrintRune(char)
	}
	ft.PrintRune('\n')
}
