/*
2064 : https://leetcode.com/problems/minimized-maximum-of-products-distributed-to-any-store/description
1. Binary Search repeatedly with sort.Search
*/
package main

import "sort"

func minimizedMaximum(n int, quantities []int) int {
	return sort.Search(1e5, func(max int) bool {
		count := 0
		for _, q := range quantities {
			count += (q + max) / (max + 1)
		}
		return count <= n
	}) + 1
}

func main() {

}
