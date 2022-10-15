package main

import (
	"fmt"
	"time"
)

func main() {
	sundays := 0
	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			t := time.Date(year, time.Month(month), 1, 1, 0, 0, 0, time.UTC)
			if t.Weekday() == 0 {
				sundays++
			}
		}
	}
	fmt.Printf("%d\n", sundays)
}
