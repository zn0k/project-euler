package main

import "fmt"

var small map[int]string = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var tens map[int]string = map[int]string{
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func translate100s(n int) (string, int) {
	if n < 100 {
		return "", n
	}
	h := n / 100
	result := small[h] + "hundred"
	n = n % 100
	return result, n
}

func translate10s(n int) string {
	if n < 20 {
		return small[n]
	}
	d := n / 10
	result := tens[d*10]
	m := n % 10
	if m != 0 {
		result += small[m]
	}
	return result
}

func translate(n int) string {
	if n == 1000 {
		return "onethousand"
	}
	hundreds, remainder := translate100s(n)
	tens := translate10s(remainder)
	if len(hundreds) > 0 && len(tens) > 0 {
		return hundreds + "and" + tens
	} else if len(hundreds) > 0 {
		return hundreds
	} else {
		return tens
	}
}

func main() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += len(translate(i))
	}
	fmt.Printf("%d\n", sum)
}
