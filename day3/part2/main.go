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

	str := recursiveHelper(line, 12)
	toAdd, err := strconv.Atoi(str)
	if err != nil {
		panic("AHHHH")
	}
	mu.Lock()
	count += toAdd
	mu.Unlock()
	wg.Done()
}

func recursiveHelper(line string, iteration int) string {
	if iteration == 0 {
		return ""
	}
	toPickFrom := line[:len(line)-iteration+1]

	max := 0
	idx := 0
	for i, r := range toPickFrom {
		num := int(r - '0')
		if num > max {
			max = num
			idx = i
		}
	}

	ret := recursiveHelper(line[idx+1:], iteration-1)

	buf := make([]byte, 0)
	buf = strconv.AppendInt(buf, int64(max), 10)
	buf = append(buf, ret...)

	return string(buf)
}
