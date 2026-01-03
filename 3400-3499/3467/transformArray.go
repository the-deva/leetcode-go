/*
3467 - https://leetcode.com/problems/transform-array-by-parity/description/
1. Find even and odd numbers in the array.
2. Replace even numbers with 0 and odd numbers with 1.
3. Sort the transformed array in non-decreasing order.
*/
package main

import "sort"

func transformArray(nums []int) []int {
	result := make([]int, 0)
	for _, num := range nums {
		if num%2 == 0 {
			result = append(result, 0)
		} else {
			result = append(result, 1)
		}
	}
	sort.Ints(result)
	return result
}

func main() {

}
