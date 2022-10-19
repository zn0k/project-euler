package main

import "fmt"

func change(n, sum int, coins []int) int {
	var i, j, x, y int

	var table [][]int
	for i = 1; i <= sum+1; i++ {
		row := make([]int, n)
		table = append(table, row)
	}

	for i = 0; i < n; i++ {
		table[0][i] = 1
	}

	for i = 1; i < sum+1; i++ {
		for j = 0; j < n; j++ {
			if i-coins[j] >= 0 {
				x = table[i-coins[j]][j]
			} else {
				x = 0
			}
			if j >= 1 {
				y = table[i][j-1]
			} else {
				y = 0
			}
			table[i][j] = x + y
		}
	}
	return table[sum][n-1]
}

func main() {
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	n := len(coins)
	sum := 200
	fmt.Printf("%d\n", change(n, sum, coins))
}
