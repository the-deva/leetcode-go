/*
2799 : https://leetcode.com/problems/maximum-beauty-of-an-array-after-applying-operation/description
1. sort the nums
2. find the longest continuous subarray whose max value - min value <= 2k
*/
package main

import "slices"

func maximumBeauty(nums []int, k int) (result int) {
	slices.Sort(nums)
	left := 0
	for right, x := range nums {
		for x-nums[left] > k*2 {
			left++
		}
		result = max(result, right-left+1)
	}
	return
}

func main() {

}
