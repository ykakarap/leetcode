package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 4, 5}
	fmt.Printf("Breaking is possible : %t \n", isPossible(nums))

}

func isPossible(nums []int) bool {
	subs := [][]int{}
	require := map[int][]int{}

	for _, v := range nums {
		// check if some sub sequence needs this number
		if ss, ok := require[v]; ok {
			// someone needs this
			// update the sequence
			// find the smallest sub and use that
			useIdx := 0
			for idx, value := range ss {
				if len(subs[value]) < len(subs[ss[useIdx]]) {
					useIdx = idx
				}
			}
			subs[ss[useIdx]] = append(subs[ss[useIdx]], v)

			// update the requirements
			if _, okay := require[v+1]; okay {
				// add to existing requirement
				require[v+1] = append(require[v+1], ss[useIdx])
			} else {
				// register a new requirement
				require[v+1] = []int{ss[useIdx]}
			}

			if len(ss) > 1 {
				require[v] = append(ss[0:useIdx], ss[useIdx+1:]...)
			} else {
				delete(require, v)
			}

		} else {
			// no one needs this number
			// start a new subsequence
			subs = append(subs, []int{v})
			// regiter the requirement
			if _, okay := require[v+1]; okay {
				// add to existing requirement
				require[v+1] = append(require[v+1], len(subs)-1)
			} else {
				// register a new requirement
				require[v+1] = []int{len(subs) - 1}
			}
		}
	}

	for _, v := range subs {
		if len(v) < 3 {
			return false
		}
	}

	return true
}
