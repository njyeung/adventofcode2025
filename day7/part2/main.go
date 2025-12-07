package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Pair struct {
	A int
	B int
}

var (
	m  = make(map[Pair]int)
	mu sync.Mutex
)
var grid [][]string

func main() {
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

	for idx, p := range grid[0] {
		if p == "S" {
			fmt.Println(recurse(1, idx))
		}
	}
}

func recurse(level int, position int) int {
	if level == len(grid) {
		return 1
	}

	mu.Lock()
	value, inMap := m[Pair{A: level, B: position}]
	mu.Unlock()

	if inMap {
		return value
	}
	for idx, c := range grid[level] {
		if c == "^" && position == idx {
			ch := make(chan int, 2)

			go func() {
				ch <- recurse(level+1, position+1)
			}()
			go func() {
				ch <- recurse(level+1, position-1)
			}()

			first := <-ch
			second := <-ch

			mu.Lock()
			m[Pair{A: level, B: position}] = first + second
			mu.Unlock()

			return first + second
		}
	}

	return recurse(level+1, position)
}
