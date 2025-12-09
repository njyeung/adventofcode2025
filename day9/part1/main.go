package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

var points []Point

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic("couldn't open file")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		coord, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("AHHH")
		}
		x := coord
		coord, err = strconv.Atoi(parts[1])
		if err != nil {
			panic("AHHH")
		}
		y := coord

		points = append(points, Point{x, y})
	}

	currMax := 0
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			size := math.Abs(float64(points[i].x-points[j].x+1)) * math.Abs(float64(points[i].y-points[j].y+1))
			if int(size) > currMax {
				currMax = int(size)
			}
		}
	}

	fmt.Println(currMax)
}
