package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	num := big.NewInt(1)

	var i int64
	for i = 1; i <= 100; i++ {
		num = num.Mul(num, big.NewInt(i))
	}

	var sum int64 = 0
	for _, c := range num.String() {
		i, err := strconv.ParseInt(string(c), 10, 64)
		if err != nil {
			panic(err)
		}
		sum += i
	}
	fmt.Printf("%d\n", sum)
}
