package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/stat/combin"
)

func digits(n int) []int {
	var result []int

	for n > 10 {
		result = append(result, n%10)
		n /= 10
	}
	result = append(result, n)
	for i := 0; i < len(result)/2; i++ {
		j := len(result) - i - 1
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func combine(s []int) int {
	r := 0
	l := len(s)
	for j, k := range s {
		r += k * int((math.Pow10(l - j - 1)))
	}
	return r
}

func check(p []int) bool {
	if p[3]%2 != 0 {
		return false
	}
	if p[5] != 5 && p[5] != 0 {
		return false
	}
	d35 := combine(p[2:5])
	if d35%3 != 0 {
		return false
	}
	d57 := combine(p[4:7])
	if d57%7 != 0 {
		return false
	}
	d68 := combine(p[5:8])
	if d68%11 != 0 {
		return false
	}
	d79 := combine(p[6:9])
	if d79%13 != 0 {
		return false
	}
	d810 := combine(p[7:10])
	if d810%17 != 0 {
		return false
	}
	return true
}

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	gen := combin.NewPermutationGenerator(len(data), len(data))
	sum := 0
	for gen.Next() {
		pan := gen.Permutation(nil)
		if check(pan) {
			sum += combine(pan)
		}
	}
	fmt.Println(sum)
}
