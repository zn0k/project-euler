package main

import (
	"fmt"
	"sort"
	"strconv"
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

func expand(n, m int) (bool, int) {
	s := ""
	for i := 1; i <= m; i++ {
		s += strconv.FormatInt(int64(n*i), 10)
	}
	if len(s) != 9 {
		return false, 0
	}
	p64, _ := strconv.ParseInt(s, 10, 64)
	p := int(p64)
	if p > 987654321 {
		return false, 0
	}

	ds := digits(p)
	sort.Ints(ds)
	for i := 1; i < 10; i++ {
		if ds[i-1] != i {
			return false, 0
		}
	}
	return true, p
}

func main() {
	max := 0
	for i := 1; i < 10000; i++ {
		for n := 2; n < 10; n++ {
			ok, pan := expand(i, n)
			if ok {
				if pan > max {
					max = pan
				}
			}
		}
	}

	fmt.Printf("%d\n", max)
}
