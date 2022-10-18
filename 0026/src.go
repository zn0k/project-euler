package main

import (
	"fmt"
	"math/big"
	"strings"
)

func cycleLength(s string, d int64) int {
	max := 0
	seg := ""
	for _, r := range s {
		segments := strings.Split(s, string(r))
		if len(segments) > 3 {
			if segments[1] == segments[2] {
				if len(segments[1]) > max {
					max = len(segments[1])
					seg = segments[1]
				}
			}
		}
	}
	fmt.Printf("1/%d has cycle length %d (%s)\n", d, max, seg)
	fmt.Println(s)
	return max
}

func main() {
	max := 0
	answer := int64(0)
	for d := int64(1); d < 1000; d++ {
		a := big.NewRat(1, d)
		_, s, _ := strings.Cut(a.FloatString(500), ".")
		l := cycleLength(s, d)
		if l > max {
			max = l
			answer = d
		}
	}
	fmt.Printf("%d\n", answer)
}
