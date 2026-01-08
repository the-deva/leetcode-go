/*
1458 - https://leetcode.com/problems/max-dot-product-of-two-subsequences/
*/
package main

import "math"

func maxDotProduct(nums1 []int, nums2 []int) int {
	m := len(nums2)
	f := make([]int, m+1)
	for i := range f {
		f[i] = math.MinInt
	}

	for _, x := range nums1 {
		pre := f[0]
		for j, y := range nums2 {
			tmp := f[j+1]
			f[j+1] = max(max(pre, 0)+x*y, f[j+1], f[j])
			pre = tmp
		}
	}
	return f[m]
}

func main() {

}
