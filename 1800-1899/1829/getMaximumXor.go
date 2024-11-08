/*
1829 :https://leetcode.com/problems/maximum-xor-for-each-query/
*/
package main

func getMaximumXor(nums []int, maximumBit int) []int {
	N := len(nums)
	result := make([]int, N)
	s, m := 0, 1<<maximumBit-1

	for i, v := range nums {
		s ^= v
		result[N-1-i] = s&m ^ m
	}
	return result
}

func main() {
	nums := []int{0, 1, 1, 3}
	maximumBit := 3
	getMaximumXor(nums, maximumBit)
}
