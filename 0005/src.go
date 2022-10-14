package main

import "fmt"

func main() {
	i := 1
	for k := 1; k <= 20; k++ {
		if i%k > 0 {
			for j := 1; j <= 20; j++ {
				if (i*j)%k == 0 {
					i *= j
					break
				}
			}
		}
	}
	fmt.Printf("%d\n", i)
}
