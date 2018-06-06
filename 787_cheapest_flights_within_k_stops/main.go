package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 3
	// flights := [][]int{
	// 	[]int{3, 4, 4},
	// 	[]int{2, 5, 6},
	// 	[]int{4, 7, 10},
	// 	[]int{9, 6, 5},
	// 	[]int{7, 4, 4},
	// 	[]int{6, 2, 10},
	// 	[]int{6, 8, 6},
	// 	[]int{7, 9, 4},
	// 	[]int{1, 5, 4},
	// 	[]int{1, 0, 4},
	// 	[]int{9, 7, 3},
	// 	[]int{7, 0, 5},
	// 	[]int{6, 5, 8},
	// 	[]int{1, 7, 6},
	// 	[]int{4, 0, 9},
	// 	[]int{5, 9, 1},
	// 	[]int{8, 7, 3},
	// 	[]int{1, 2, 6},
	// 	[]int{4, 1, 5},
	// 	[]int{5, 2, 4},
	// 	[]int{1, 9, 1},
	// 	[]int{7, 8, 10},
	// 	[]int{0, 4, 2},
	// 	[]int{7, 2, 8},
	// }

	// flights := [][]int{
	// 	[]int{0, 12, 28},
	// 	[]int{5, 6, 39},
	// 	[]int{8, 6, 59},
	// 	[]int{13, 15, 7},
	// 	[]int{13, 12, 38},
	// 	[]int{10, 12, 35},
	// 	[]int{15, 3, 23},
	// 	[]int{7, 11, 26},
	// 	[]int{9, 4, 65},
	// 	[]int{10, 2, 38},
	// 	[]int{4, 7, 7},
	// 	[]int{14, 15, 31},
	// 	[]int{2, 12, 44},
	// 	[]int{8, 10, 34},
	// 	[]int{13, 6, 29},
	// 	[]int{5, 14, 89},
	// 	[]int{11, 16, 13},
	// 	[]int{7, 3, 46},
	// 	[]int{10, 15, 19},
	// 	[]int{12, 4, 58},
	// 	[]int{13, 16, 11},
	// 	[]int{16, 4, 76},
	// 	[]int{2, 0, 12},
	// 	[]int{15, 0, 22},
	// 	[]int{16, 12, 13},
	// 	[]int{7, 1, 29},
	// 	[]int{7, 14, 100},
	// 	[]int{16, 1, 14},
	// 	[]int{9, 6, 74},
	// 	[]int{11, 1, 73},
	// 	[]int{2, 11, 60},
	// 	[]int{10, 11, 85},
	// 	[]int{2, 5, 49},
	// 	[]int{3, 4, 17},
	// 	[]int{4, 9, 77},
	// 	[]int{16, 3, 47},
	// 	[]int{15, 6, 78},
	// 	[]int{14, 1, 90},
	// 	[]int{10, 5, 95},
	// 	[]int{1, 11, 30},
	// 	[]int{11, 0, 37},
	// 	[]int{10, 4, 86},
	// 	[]int{0, 8, 57},
	// 	[]int{6, 14, 68},
	// 	[]int{16, 8, 3},
	// 	[]int{13, 0, 65},
	// 	[]int{2, 13, 6},
	// 	[]int{5, 13, 5},
	// 	[]int{8, 11, 31},
	// 	[]int{6, 10, 20},
	// 	[]int{6, 2, 33},
	// 	[]int{9, 1, 3},
	// 	[]int{14, 9, 58},
	// 	[]int{12, 3, 19},
	// 	[]int{11, 2, 74},
	// 	[]int{12, 14, 48},
	// 	[]int{16, 11, 100},
	// 	[]int{3, 12, 38},
	// 	[]int{12, 13, 77},
	// 	[]int{10, 9, 99},
	// 	[]int{15, 13, 98},
	// 	[]int{15, 12, 71},
	// 	[]int{1, 4, 28},
	// 	[]int{7, 0, 83},
	// 	[]int{3, 5, 100},
	// 	[]int{8, 9, 14},
	// 	[]int{15, 11, 57},
	// 	[]int{3, 6, 65},
	// 	[]int{1, 3, 45},
	// 	[]int{14, 7, 74},
	// 	[]int{2, 10, 39},
	// 	[]int{4, 8, 73},
	// 	[]int{13, 5, 77},
	// 	[]int{10, 0, 43},
	// 	[]int{12, 9, 92},
	// 	[]int{8, 2, 26},
	// 	[]int{1, 7, 7},
	// 	[]int{9, 12, 10},
	// 	[]int{13, 11, 64},
	// 	[]int{8, 13, 80},
	// 	[]int{6, 12, 74},
	// 	[]int{9, 7, 35},
	// 	[]int{0, 15, 48},
	// 	[]int{3, 7, 87},
	// 	[]int{16, 9, 42},
	// 	[]int{5, 16, 64},
	// 	[]int{4, 5, 65},
	// 	[]int{15, 14, 70},
	// 	[]int{12, 0, 13},
	// 	[]int{16, 14, 52},
	// 	[]int{3, 10, 80},
	// 	[]int{14, 11, 85},
	// 	[]int{15, 2, 77},
	// 	[]int{4, 11, 19},
	// 	[]int{2, 7, 49},
	// 	[]int{10, 7, 78},
	// 	[]int{14, 6, 84},
	// 	[]int{13, 7, 50},
	// 	[]int{11, 6, 75},
	// 	[]int{5, 10, 46},
	// 	[]int{13, 8, 43},
	// 	[]int{9, 10, 49},
	// 	[]int{7, 12, 64},
	// 	[]int{0, 10, 76},
	// 	[]int{5, 9, 77},
	// 	[]int{8, 3, 28},
	// 	[]int{11, 9, 28},
	// 	[]int{12, 16, 87},
	// 	[]int{12, 6, 24},
	// 	[]int{9, 15, 94},
	// 	[]int{5, 7, 77},
	// 	[]int{4, 10, 18},
	// 	[]int{7, 2, 11},
	// 	[]int{9, 5, 41},
	// }

	flights := [][]int{
		[]int{0, 1, 100},
		[]int{1, 0, 100},
		[]int{1, 2, 100},
		[]int{0, 2, 500},
	}
	src := 0
	dst := 2
	k := 1
	fmt.Printf("Answer : %v \n", findCheapestPrice(n, flights, src, dst, k))
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	graph := map[int][][]int{}
	// construct the map
	for _, v := range flights {
		if _, ok := graph[v[0]]; ok {
			graph[v[0]] = append(graph[v[0]], []int{v[1], v[2]})
		} else {
			graph[v[0]] = [][]int{}
			graph[v[0]] = append(graph[v[0]], []int{v[1], v[2]})
		}
	}

	r := &Routes{}
	f := Flight{}
	findAllRoutes(graph, src, dst, map[int]bool{}, f, r)

	sort.Slice(r.Flights, func(i, j int) bool {
		return r.Flights[i].Cost < r.Flights[j].Cost
	})

	fmt.Println(r.Flights)

	// get the flight within stops limit
	for _, f := range r.Flights {
		if f.Cities-1 <= K {
			return f.Cost
		}
	}

	return -1
}

type Routes struct {
	Flights []Flight
}

type Flight struct {
	Path   []int
	Cost   int
	Cities int // We actually get an extra stop here
}

func findAllRoutes(graph map[int][][]int, src, dst int, visited map[int]bool, f Flight, r *Routes) {
	if v, ok := visited[src]; ok && v {
		// possibe cyclic
		return
	}

	// add current location to path and mark as visited
	f.Path = append(f.Path, src)
	visited[src] = true

	// check if we are at destination
	if src == dst {
		// we have reached destination
		// store the path
		tmp := make([]int, len(f.Path))
		copy(tmp, f.Path)
		tempF := Flight{
			Path:   tmp,
			Cities: f.Cities,
			Cost:   f.Cost,
		}
		r.Flights = append(r.Flights, tempF)

	}

	// we are not yet at destination
	// check all connecting routes
	for _, v := range graph[src] {
		tmpF := Flight{
			Path:   f.Path,
			Cities: f.Cities + 1,
			Cost:   f.Cost + v[1],
		}
		findAllRoutes(graph, v[0], dst, visited, tmpF, r)
	}

	// done with this path
	// reset visited
	visited[src] = false
}
