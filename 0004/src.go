package main

import "fmt"

func palindrome(num int) bool {
	var digits []int
	for num > 10 {
		digits = append(digits, num%10)
		num /= 10
	}
	digits = append(digits, num)
	l := len(digits) - 1
	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[l-i] {
			return false
		}
	}
	return true
}

func main() {
	max := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			p := i * j
			if palindrome(p) {
				if p > max {
					max = p
				}
			}
		}
	}
	fmt.Printf("%d\n", max)
}
