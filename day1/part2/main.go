package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	count := 0
	curr := 50
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		toTurn, err := strconv.Atoi(line[1:])
		if err != nil {
			panic("could not convert to int")
		}

		count += toTurn / 100
		toTurn = toTurn % 100

		if line[0] == 'L' {
			prev := curr
			curr -= toTurn
			if curr <= 0 && prev != 0 {
				count += 1
			}
			if curr < 0 {
				curr = 100 + curr
			}
		} else { // R
			prev := curr
			curr += toTurn
			if curr > 99 {
				if prev != 0 {
					count += 1
				}
				curr = curr - 100
			}
		}

		fmt.Println(curr)
		fmt.Printf("COUNT: %d\n", count)
		fmt.Println()
	}

	fmt.Println(count)
}
