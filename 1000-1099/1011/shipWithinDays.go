/*
1011 : https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/description/
1. Binary search
2. Iterate over weights and if w + pre > num then cost++
*/
package main

import (
	"leetcode-go/utils"
)

func shipWithinDays(weights []int, days int) int {
	low := utils.Max(weights)  // set to max of weights
	high := utils.Sum(weights) // set to sum of weights

	// Binary search the solution
	check := func(num int) bool {
		pre := 0
		cost := 0

		for _, w := range weights {
			if w+pre > num {
				cost++
				pre = 0
			}
			pre += w
		}
		return cost+1 <= days // cost will be atleast 1
	}

	for low < high-1 {
		mid := low + (high-low)/2
		if check(mid) {
			high = mid
		} else {
			low = mid
		}
	}
	if check(low) {
		return low
	}
	return high
}

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days := 5
	shipWithinDays(weights, days)
}
