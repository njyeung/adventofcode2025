package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var grid [][]string

	file, err := os.Open("input.txt")
	if err != nil {
		panic("couldn't open file")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")
		grid = append(grid, parts)
	}

	m := len(grid)
	n := len(grid[0])

	count := 0
	curr := 0
	op := "+"
	for row := 0; row < n; row++ {
		buff := make([]string, m)

		for col := 0; col < m; col++ {
			buff[col] = grid[col][row]
		}

		if buff[m-1] == "+" {
			count += curr
			curr = 0
			op = "+"
		}
		if buff[m-1] == "*" {
			count += curr
			curr = 1
			op = "*"
		}

		digits := make([]rune, 0)
		for _, s := range buff {
			for _, r := range s {
				if unicode.IsDigit(r) {
					digits = append(digits, r)
				}
			}

		}
		if len(digits) == 0 {
			continue
		}
		num, err := strconv.Atoi(string(digits))
		if err != nil {
			panic("AHHHHHHH")
		}

		if op == "+" {
			curr += num
		} else {
			curr *= num
		}
	}
	count += curr

	fmt.Println(count)
}
