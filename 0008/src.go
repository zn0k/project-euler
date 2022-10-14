package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	var digits []int
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		for _, c := range line {
			i, err := strconv.ParseInt(string(c), 10, 32)
			if err != nil {
				panic(err)
			}
			digits = append(digits, int(i))
		}
	}

	max := 0

	for i := 0; i < len(digits)-13; i++ {
		p := 1
		for j := 0; j < 13; j++ {
			p *= digits[i+j]
		}
		if p > max {
			max = p
		}
	}

	fmt.Printf("%d\n", max)
}
