package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(p string) []rune {
	f, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var result []rune
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		for _, c := range fields {
			i, err := strconv.ParseInt(c, 10, 32)
			if err != nil {
				panic(err)
			}
			result = append(result, int32(i))
		}
	}
	return result
}

func decode(input []rune, cipher []rune) []rune {
	l := len(cipher)
	rs := make([]rune, len(input))
	for c := 0; c < len(cipher); c++ {
		for i := c; i < len(input); i += l {
			rs[i] = input[i] ^ cipher[c]
		}
	}
	return rs
}

func toRunes(s string) []rune {
	var result []rune
	for _, c := range s {
		result = append(result, rune(c))
	}
	return result
}

func toString(rs []rune) string {
	s := ""
	for _, r := range rs {
		s += string(r)
	}
	return s
}

func main() {
	input := readLines("input.txt")
	sum := 0
loop:
	for i := 'a'; i <= 'z'; i++ {
		for j := 'a'; j <= 'z'; j++ {
			for k := 'a'; k <= 'z'; k++ {
				cipher := []rune{i, j, k}
				decoded := decode(input, cipher)
				s := toString(decoded)
				if strings.Contains(s, "the") && strings.Contains(s, " a ") {
					for _, r := range decoded {
						sum += int(r)
					}
					break loop
				}
			}
		}
	}
	fmt.Println(sum)
}
