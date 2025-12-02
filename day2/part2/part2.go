package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	count int = 0
	mu    sync.Mutex
	wg    sync.WaitGroup
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	parts := strings.Split(line, ",")
	for _, part := range parts {

		dist := strings.Split(part, "-")
		first, err := strconv.Atoi(dist[0])
		if err != nil {
			panic("parsing err")
		}
		second, err := strconv.Atoi(dist[1])
		if err != nil {
			panic("parsing err")
		}

		wg.Add(1)
		go countInvalid(first, second)
	}
	wg.Wait()
	fmt.Println(count)
}

func countInvalid(start int, end int) {
	invalidCount := 0
	for i := start; i <= end; i++ {
		num := strconv.FormatInt(int64(i), 10)
		if isInvalid(num) {
			invalidCount += i
		}
	}
	mu.Lock()
	count += invalidCount
	mu.Unlock()
	wg.Done()
}

func isInvalid(number string) bool {
	for i := 1; i <= len(number)/2; i++ {
		if len(number)%i != 0 {
			continue
		}

		root := number[:i]
		isValid := true
		for j := i; j < len(number); j += i {
			if root != number[j:j+i] {
				isValid = false
				break
			}
		}
		if isValid {
			return true
		}
	}
	return false
}
