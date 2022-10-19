package main

import "fmt"

type triangle struct {
	a, b, c int
}

type candidate struct {
	p, solutions int
}

func solve(p int) candidate {
	var solutions []triangle
	for a := 1; a <= p; a++ {
		for b := 1; b <= p-a; b++ {
			for c := 1; c <= p-a-b; c++ {
				if a+b+c == p {
					if a*a+b*b == c*c {
						solutions = append(solutions, triangle{a, b, c})
					}
				}
			}
		}
	}
	return candidate{p, len(solutions) / 2}
}

func main() {
	max := candidate{0, 0}
	for p := 1; p <= 1000; p++ {
		c := solve(p)
		if c.solutions > max.solutions {
			max = c
		}
	}
	fmt.Printf("%d\n", max.p)
}
