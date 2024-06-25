package piscine

import "ft"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	nums :=  make([]rune, n)
	first := true
	printCombNRecursive(nums, 0, '0', &first)
	ft.PrintRune('\n')
}

func printCombNRecursive(nums []rune, index int, startChar rune, first *bool) {
	if index == len(nums) {
		if *first {
			*first = false
		} else {
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
		for _, num := range nums {
			ft.PrintRune(num)
		}
		return
	}

	for i := startChar; i <= '9'; i++ {
		nums[index] = i
		printCombNRecursive(nums, index+1, i+1, first)
	}
}

