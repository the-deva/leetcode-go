/*
2924 : https://leetcode.com/problems/find-champion-ii/description
1. Mark all the edges[1] to true as they are not champions
2. Iterate over the weak and update result and return -1 if result has to be updated again.
*/
package main

func findChampion(n int, edges [][]int) int {
	weak := make([]bool, n)
	for _, e := range edges {
		weak[e[1]] = true // Marking all e[1] to true as they are not champions
	}

	result := -1
	for i, w := range weak {
		if w {
			continue
		}
		if result != -1 {
			return -1 // There can be only champion
		}
		result = i
	}
	return result
}

func main() {

}
