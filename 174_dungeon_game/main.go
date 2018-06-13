package main

import (
	"fmt"
)

func main() {

	dungeon := [][]int{
		[]int{1, -3, 3},
		[]int{0, -2, 0},
		[]int{-3, -3, -3},
	}
	// dungeon := [][]int{
	// 	[]int{-2, -3, 3},
	// 	[]int{-5, -10, 1},
	// 	[]int{10, 30, -5},
	// }

	fmt.Printf("Minimum energy required by the Knight : %v \n", calculateMinimumHP(dungeon))

}

func calculateMinimumHP(dungeon [][]int) int {
	rows := len(dungeon)
	cols := len(dungeon[0])

	requirement := make([]int, rows*cols)
	// for i := 0; i < rows; i++ {
	// 	requirement[i] = make([]int, cols)
	// }

	for r := rows - 1; r >= 0; r-- {
		for c := cols - 1; c >= 0; c-- {
			var rightExists, bottomExits bool
			if c < cols-1 {
				rightExists = true
			}
			if r < rows-1 {
				bottomExits = true
			}

			var req int
			if rightExists && bottomExits {
				// req = min(requirement[r+1][c], requirement[r][c+1]) - dungeon[r][c]
				req = min(requirement[(r+1)*cols+c], requirement[r*cols+c+1]) - dungeon[r][c]
			} else if rightExists {
				req = requirement[r*cols+c+1] - dungeon[r][c]
			} else if bottomExits {
				req = requirement[(r+1)*cols+c] - dungeon[r][c]
			} else {
				req = 0 - dungeon[r][c]
			}

			if req <= 0 {
				requirement[r*cols+c] = 0
			} else {
				requirement[r*cols+c] = req
			}
		}
	}

	return requirement[0] + 1
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
