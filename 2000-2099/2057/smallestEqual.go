/*
2057 : https://leetcode.com/problems/smallest-index-with-equal-value/description/
*/
package main

func smallestEqual(nums []int) int {
	for i, num := range nums {
		if i%10 == num {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{0, 1, 2}
	smallestEqual(nums)
}
