package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var (
	count int
	mu    sync.Mutex
	wg    sync.WaitGroup
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("couldn't open file")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go solveLine(line)
	}
	wg.Wait()
	fmt.Println(count)
}

func solveLine(line string) {
	first := 0
	first_idx := -1
	second := 0

	r := []rune(line)
	for i := 0; i < len(r)-1; i++ {
		ru := line[i]
		num := int(ru - '0')

		if num > first {
			first = num
			first_idx = i
		}
	}

	for i := first_idx + 1; i < len(r); i++ {
		ru := line[i]
		num := int(ru - '0')

		if num > second {
			second = num
		}
	}

	buf := make([]byte, 0, 4)
	buf = strconv.AppendInt(buf, int64(first), 10)
	buf = strconv.AppendInt(buf, int64(second), 10)

	toAdd, err := strconv.Atoi(string(buf))
	if err != nil {
		panic("could not convert to int")
	}
	mu.Lock()
	count += toAdd
	mu.Unlock()
	wg.Done()
}
