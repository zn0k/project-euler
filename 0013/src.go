package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	nums := make([]*big.Int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		num := new(big.Int)
		num, ok := num.SetString(line, 10)
		if ok {
			nums = append(nums, num)
		}
	}

	sum := new(big.Int)
	for _, i := range nums {
		sum.Add(sum, i)
	}

	for i, c := range sum.String() {
		if i >= 10 {
			break
		}
		fmt.Printf("%s", string(c))
	}
	fmt.Printf("\n")
}
