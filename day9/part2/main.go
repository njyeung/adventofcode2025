package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type PointSet map[Point]struct{}

func (s PointSet) Add(p Point) {
	s[p] = struct{}{}
}

func (s PointSet) Remove(p Point) {
	delete(s, p)
}

func (s PointSet) Has(p Point) bool {
	_, ok := s[p]
	return ok
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("couldn't open file")
	}
	defer file.Close()

	var polyReal []Point

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		polyReal = append(polyReal, Point{x, y})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// get logical polygon outline
	polyLogical, grid, logicalToReal := buildLogicalGrid(polyReal)

	// fill in the polygon
	fillPolygon(grid)

	bestArea := 0
	for i := 0; i < len(polyLogical); i++ {
		for j := i + 1; j < len(polyLogical); j++ {
			a := polyLogical[i]
			b := polyLogical[j]

			// logical bounding box
			x1 := a.x
			x2 := b.x
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			y1 := a.y
			y2 := b.y
			if y1 > y2 {
				y1, y2 = y2, y1
			}

			// goofy ahh brute force
			if !rectFullyInside(grid, x1, y1, x2, y2) {
				continue
			}

			// convert back to real coords
			dx := abs(logicalToReal[a].x-logicalToReal[b].x) + 1
			dy := abs(logicalToReal[a].y-logicalToReal[b].y) + 1

			bestArea = max(bestArea, dx*dy)
		}
	}

	fmt.Println(bestArea)
}

func buildLogicalGrid(polyReal []Point) (polyLogcal []Point, grid [][]int, logicalToReal map[Point]Point) {
	logicalToReal = make(map[Point]Point, len(polyReal))

	// this is disgusting
	var sortedByX []Point
	var sortedByY []Point
	for _, p := range polyReal {
		sortedByX = append(sortedByX, p)
		sortedByY = append(sortedByY, p)
	}
	sort.Slice(sortedByX, func(i, j int) bool {
		return sortedByX[i].x < sortedByX[j].x
	})
	sort.Slice(sortedByY, func(i, j int) bool {
		return sortedByY[i].y < sortedByY[j].y
	})

	var x []PointSet
	prev := -1
	var buff PointSet
	for _, p := range sortedByX {
		if p.x == prev {
			buff.Add(p)
		} else {
			x = append(x, buff)
			buff = PointSet{}
			buff.Add(p)
			prev = p.x
		}
	}
	x = append(x, buff)
	x = x[1:]

	var y []PointSet
	prev = -1
	buff = PointSet{}
	for _, p := range sortedByY {
		if p.y == prev {
			buff.Add(p)
		} else {
			y = append(y, buff)
			buff = PointSet{}
			buff.Add(p)
			prev = p.y
		}
	}
	y = append(y, buff)
	y = y[1:]

	W := len(polyReal)/2 + 2
	H := len(polyReal)/2 + 2
	grid = make([][]int, W)
	for i := range grid {
		grid[i] = make([]int, H)
	}

	// Get points
	var polyLogical []Point
	for _, p := range polyReal {
		xIdx := -1
		for idx, ps := range x {
			if ps.Has(p) {
				xIdx = idx + 1
				break
			}
		}
		yIdx := -1
		for idx, ps := range y {
			if ps.Has(p) {
				yIdx = idx + 1
				break
			}
		}
		polyLogical = append(polyLogical, Point{xIdx, yIdx})
		grid[xIdx][yIdx] = 1
		logicalToReal[Point{xIdx, yIdx}] = p
	}

	// draw line between points
	for j := 0; j < len(polyLogical); j++ {
		i := j - 1
		if i < 0 {
			i = len(polyLogical) - 1
		}

		// guaranteed to be next to each other
		a := polyLogical[i]
		b := polyLogical[j]

		// vertical edge
		if a.x == b.x {
			x := a.x
			y1, y2 := a.y, b.y
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				grid[x][y] = 1
			}
		} else if a.y == b.y { // horizontal edge
			y := a.y
			x1, x2 := a.x, b.x
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for x := x1; x <= x2; x++ {
				grid[x][y] = 1
			}
		} else {
			panic("non axis-aligned edge between points")
		}
	}

	return polyLogical, grid, logicalToReal
}

func fillPolygon(grid [][]int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return
	}

	W := len(grid)
	H := len(grid[0])

	startX, startY := W-1, H-1

	// BFS queue
	queue := []Point{{startX, startY}}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// Flood-fill
	// mark all reachable 0's as -1 (outside)
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		x, y := c.x, c.y

		if grid[x][y] != 0 {
			continue
		}

		grid[x][y] = -1

		for _, d := range dirs {
			newX, newY := x+d[0], y+d[1]
			if newX < 0 || newX >= W || newY < 0 || newY >= H {
				continue
			}
			if grid[newX][newY] == 0 {
				queue = append(queue, Point{newX, newY})
			}
		}
	}

	// -1 = outside
	// 0 = inside
	// 1 = boundary
	// In final representation, clean up with 1 on inside and 0 on outside
	for x := 0; x < W; x++ {
		for y := 0; y < H; y++ {
			switch grid[x][y] {
			case 0:
				// inside
				grid[x][y] = 1
			case -1:
				// outside
				grid[x][y] = 0
			}
		}
	}
}

func rectFullyInside(grid [][]int, x1, y1, x2, y2 int) bool {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if grid[x][y] != 1 {
				return false
			}
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
