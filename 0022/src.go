package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func scoreWord(name string) int {
	score := 0
	for _, r := range name {
		score += int(r) - 64
	}
	return score
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var names []string

	for scanner.Scan() {
		for _, f := range strings.Split(scanner.Text(), ",") {
			name := strings.ReplaceAll(f, "\"", "")
			names = append(names, name)
		}
	}
	sort.Strings(names)

	sum := 0
	for i, name := range names {
		sum += ((i + 1) * scoreWord(name))
	}
	fmt.Printf("%d\n", sum)
}
