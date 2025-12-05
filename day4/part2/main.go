package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid struct {
	Cells [][]rune
	Rows  int
	Cols  int
}

var dirs = [][]int{
	{-1, -1}, // up left
	{-1, 0},  // up
	{-1, 1},  // up right
	{0, 1},   // right
	{1, 1},   // down right
	{1, 0},   // down
	{1, -1},  // down left
	{0, -1},  // left
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var cells [][]rune
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		cells = append(cells, []rune(scanner.Text()))
	}

	grid := Grid{
		Cells: cells,
		Rows:  len(cells),
		Cols:  len(cells[0]),
	}

	count := 0
	removed := -1
	for removed != 0 {
		removed = grid.simulate()
		count += removed
	}

	fmt.Println(count)
}

func (g *Grid) simulate() (removed int) {
	grid := g.Cells
	count := 0

	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == '@' {
				// false = there isn't a roll
				// true = there is a roll
				// up left, up, up right, right, bottom right, bottom, bottom left, left,
				mask := make([]bool, 8)

				for idx, dirs := range dirs {
					ni := i + dirs[0]
					nj := j + dirs[1]

					if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) {
						if grid[ni][nj] == '@' {
							mask[idx] = true
						}
					}
				}

				rolls := 0
				for _, i := range mask {
					if i == true {
						rolls += 1
					}
				}

				if rolls < 4 {
					count += 1
					grid[i][j] = '.'
				}
			}
		}
	}

	return count
}
