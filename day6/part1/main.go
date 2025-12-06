package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		parts := strings.Fields(line)
		grid = append(grid, parts)
	}

	count := 0
	for i := 0; i < len(grid[0]); i++ {
		buff := make([]int, len(grid)-1)
		for j := 0; j < len(grid); j++ {
			num, err := strconv.Atoi(grid[j][i])
			if err != nil {
				op := grid[j][i]
				curr := 1
				if op == "+" {
					curr = 0
				}
				for _, num := range buff {
					if op == "+" {
						curr += num
					} else {
						curr *= num
					}
				}
				count += curr
			} else {
				buff[j] = num
			}
		}
	}

	fmt.Println(count)
}
