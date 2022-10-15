package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var t [][]int

	for scanner.Scan() {
		var row []int
		for _, f := range strings.Fields(scanner.Text()) {
			i, err := strconv.ParseInt(f, 10, 32)
			if err != nil {
				panic(err)
			}
			row = append(row, int(i))
		}
		t = append(t, row)
	}

	for i := len(t) - 2; i >= 0; i-- {
		for j := 0; j < len(t[i]); j++ {
			if t[i+1][j] > t[i+1][j+1] {
				t[i][j] += t[i+1][j]
			} else {
				t[i][j] += t[i+1][j+1]
			}
		}
	}
	fmt.Printf("%d\n", t[0][0])
}
