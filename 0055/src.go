package main

import (
	"fmt"
	"math/big"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		j := len(runes) - 1 - i
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		j := len(runes) - 1 - i
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func Step(a *big.Int) {
	var b big.Int
	b.SetString(Reverse(a.String()), 10)
	a.Add(a, &b)
}

func main() {
	count := 0
	for i := 0; i < 10_000; i++ {
		num := big.NewInt(int64(i))
		isLychrel := true
		for j := 0; j < 49; j++ {
			Step(num)
			if IsPalindrome(num.String()) {
				isLychrel = false
				break
			}
		}
		if isLychrel {
			count++
		}
	}
	fmt.Println(count)
}
