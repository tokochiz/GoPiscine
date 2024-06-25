package piscine

import "ft"

func printTwoDigitNumber(nb int) {
	tens := '0' + rune(nb / 10)
	ones := '0' + rune(nb % 10)
	ft.PrintRune(tens)
	ft.PrintRune(ones)
}

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			printTwoDigitNumber(i)
			ft.PrintRune(' ')
			printTwoDigitNumber(j)
			if i == 98 && j == 99 {
				ft.PrintRune('\n')
			} else {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}	
}