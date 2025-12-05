package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	var processed []Interval
	var curr = &Interval{
		start: intervals[0].start,
		end:   intervals[0].end,
	}
	for i := 1; i < len(intervals); i++ {

		if curr.end >= intervals[i].start {
			if intervals[i].end > curr.end {
				curr.end = intervals[i].end
			}
		} else {
			processed = append(processed, *curr)
			curr = &Interval{
				start: intervals[i].start,
				end:   intervals[i].end,
			}
		}
	}

	if curr != nil {
		processed = append(processed, *curr)
	}

	fmt.Println(processed)
	count := 0
	for _, interval := range processed {
		count += interval.end - interval.start + 1
	}

	fmt.Println(count)
}
