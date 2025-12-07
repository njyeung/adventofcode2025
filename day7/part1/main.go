package main

import (
	"bufio"
	"fmt"
	"os"
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
		parts := strings.Split(line, "")
		grid = append(grid, parts)
	}

	// fmt.Println(grid)
	buff := make([]bool, len(grid[0]))

	count := 0
	for _, row := range grid {
		for idx, s := range row {
			if s == "S" {
				buff[idx] = true
			}
			if s == "^" && buff[idx] == true {
				buff[idx] = false
				buff[idx-1] = true
				buff[idx+1] = true
				count += 1
			}
		}
	}
	fmt.Println(count)
}
