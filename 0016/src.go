package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	a := big.NewInt(1)
	a.Lsh(a, 1000)

	sum := int64(0)
	for _, d := range a.String() {
		i, _ := strconv.ParseInt(string(d), 10, 64)
		sum += i
	}
	fmt.Printf("%d\n", sum)
}
