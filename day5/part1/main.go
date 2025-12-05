package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Interval struct {
	start int
	end   int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var intervals []Interval
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		seg := strings.Split(text, "-")
		start, err := strconv.Atoi(seg[0])
		if err != nil {
			panic("i cant parse")
		}
		end, err := strconv.Atoi(seg[1])
		if err != nil {
			panic("i cant parse")
		}

		intervals = append(intervals, Interval{start, end})
	}

	count := 0
	for scanner.Scan() {
		text := scanner.Text()

		ingredient, err := strconv.Atoi(text)
		if err != nil {
			panic("i cant parse")
		}

		isFresh := false
		for _, interval := range intervals {
			if ingredient >= interval.start && ingredient <= interval.end {
				isFresh = true
				break
			}
		}

		if isFresh {
			count += 1
		}
	}

	fmt.Println(count)
}
