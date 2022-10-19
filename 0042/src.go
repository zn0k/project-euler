package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func score(s string) int {
	score := 0
	for _, r := range s {
		score += int(r) - 64
	}
	return score
}

func triangleNumbers(max int) []int {
	var result []int
	num := 0
	for i := 1; num <= max; i++ {
		num = (i * (i + 1)) / 2
		result = append(result, num)
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var input []string
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), "\"", "")
		input = strings.Split(line, ",")
	}

	words := make(map[string]int)
	max := 0
	for _, w := range input {
		s := score(w)
		words[w] = s
		if s > max {
			max = s
		}
	}

	ts := make(map[int]bool)
	for _, t := range triangleNumbers(max) {
		ts[t] = true
	}

	count := 0
	for _, score := range words {
		_, ok := ts[score]
		if ok {
			count++
		}
	}
	fmt.Println(count)
}
