package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Box struct {
	index int
	x     int
	y     int
	z     int
}
type Edge struct {
	dist float64
	b1   Box
	b2   Box
}

type Set map[string]struct{}

var boxes []Box
var junctions []Set

// define heap.Interface
type MinHeap []Edge

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Edge)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("couldn't open file")
	}

	h := &MinHeap{}
	heap.Init(h)

	scanner := bufio.NewScanner(file)
	idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		box := Box{index: idx}
		coord, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("AHHH")
		}
		box.x = coord
		coord, err = strconv.Atoi(parts[1])
		if err != nil {
			panic("AHHH")
		}
		box.y = coord
		coord, err = strconv.Atoi(parts[2])
		if err != nil {
			panic("AHHH")
		}
		box.z = coord

		boxes = append(boxes, box)

		idx++
	}

	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			b1 := boxes[i]
			b2 := boxes[j]

			dist := calculateDistance(b1, b2)
			heap.Push(h, Edge{dist, b1, b2})
		}
	}

	for {
		if h.Len() == 0 {
			break
		}
		edge := heap.Pop(h).(Edge)

		b1Key := getKey(edge.b1)
		b2Key := getKey(edge.b2)

		b1JunctionIdx := -1
		b2JunctionIdx := -1
		for idx, junction := range junctions {
			if junction.Has(b1Key) {
				b1JunctionIdx = idx
			}
			if junction.Has(b2Key) {
				b2JunctionIdx = idx
			}
		}

		// cases:
		// b1 not in a junction, b2 not in a junction, create a new junction
		// b1 and b2 in the same junction, do nothing
		// b1 in a junction, b2 in a separate junction, merge the junctions
		// b1 in a junction, b2 not in a junction, add b2 to b1's junction
		// b1 not in a junction, b2 in a junction, add b1 to b2's junction
		if b1JunctionIdx == -1 && b2JunctionIdx == -1 {
			s := Set{}
			s.Add(b1Key)
			s.Add(b2Key)
			junctions = append(junctions, s)
		} else if b1JunctionIdx == b2JunctionIdx {
			// do nothing
		} else if b1JunctionIdx != -1 && b2JunctionIdx != -1 && b1JunctionIdx != b2JunctionIdx {
			i, j := b1JunctionIdx, b2JunctionIdx
			if i > j {
				i, j = j, i
			}
			s := MergeSets(junctions[i], junctions[j])
			junctions = append(junctions[:j], junctions[j+1:]...)
			junctions = append(junctions[:i], junctions[i+1:]...)
			junctions = append(junctions, s)
		} else if b1JunctionIdx != -1 && b2JunctionIdx == -1 {
			junctions[b1JunctionIdx].Add(b2Key)
		} else {
			junctions[b2JunctionIdx].Add(b1Key)
		}

		if len(junctions) == 1 && len(junctions[0]) == len(boxes) {
			fmt.Println(edge.b2.x * edge.b1.x)
			break
		}
	}
}

func (s Set) Add(v string) {
	s[v] = struct{}{}
}
func (s Set) Has(v string) bool {
	_, ok := s[v]
	return ok
}
func (s Set) Delete(v string) {
	delete(s, v)
}
func MergeSets(a, b Set) Set {
	result := make(Set)
	for k := range a {
		result[k] = struct{}{}
	}
	for k := range b {
		result[k] = struct{}{}
	}

	return result
}
func getKey(box Box) string {
	s1 := strconv.FormatInt(int64(box.x), 10)
	s2 := strconv.FormatInt(int64(box.y), 10)
	s3 := strconv.FormatInt(int64(box.z), 10)

	return s1 + s2 + s3
}
func calculateDistance(b1 Box, b2 Box) float64 {
	dx := b1.x - b2.x
	dy := b1.y - b2.y
	dz := b1.z - b2.z
	dist := math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
	return dist
}
