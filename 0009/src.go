package main

import "fmt"

func main() {
	for c := 2; c < 1000; c++ {
		for b := 2; b < c; b++ {
			for a := 2; a < b; a++ {
				if a+b+c == 1000 {
					if (a*a)+(b*b) == (c * c) {
						fmt.Printf("%d\n", a*b*c)
					}
				}
			}
		}
	}
}
