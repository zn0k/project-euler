package main

import (
	"fmt"
	"strconv"
)

func fraction() chan int {
	c := make(chan int)
	var num int64 = 1
	go func() {
		for {
			s := strconv.FormatInt(num, 10)
			for _, r := range s {
				i, _ := strconv.ParseInt(string(r), 10, 32)
				c <- int(i)
			}
			num++
		}
	}()
	return c
}

func main() {
	trigger := 1
	product := 1
	count := 1
	for digit := range fraction() {
		if count == trigger {
			product *= digit
			if count == 1_000_000 {
				break
			}
			trigger *= 10
		}
		count++
	}
	fmt.Printf("%d\n", product)
}
