package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func quickCheck(i, j int) bool {
	digits := 0
	p := i * j
	for i > 10 {
		digits++
		i /= 10
	}
	digits++
	for j > 10 {
		digits++
		j /= 10
	}
	digits++
	for p > 10 {
		digits++
		p /= 10
	}
	digits++
	return digits == 9
}

func sortString(s string) string {
	var list []string
	for _, r := range s {
		list = append(list, string(r))
	}
	sort.Strings(list)
	return strings.Join(list, "")
}

func check(i, j int) bool {
	if !quickCheck(i, j) {
		return false
	}
	s := strconv.FormatInt(int64(i), 10) + strconv.FormatInt(int64(j), 10) + strconv.FormatInt(int64(i*j), 10)
	s = sortString(s)
	return s == "123456789"
}

func main() {
	products := make(map[int]bool)
	for i := 1; i <= 10000; i++ {
		for j := 1; j <= 10000; j++ {
			if check(i, j) {
				products[i*j] = true
			}
		}
	}
	sum := 0
	for p, _ := range products {
		sum += p
	}
	fmt.Printf("%d\n", sum)
}
