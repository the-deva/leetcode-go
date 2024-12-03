/*
485 : https://leetcode.com/problems/max-consecutive-ones/description/
1. Iterate if x == 1 increment count else reset count to 0
2. Calculate max(result, count) and return
*/
package main

func findMaxConsecutiveOnes(nums []int) (result int) {
	count := 0
	for _, x := range nums {
		if x == 1 {
			count++
		} else {
			count = 0
		}
		result = max(result, count)
	}
	return
}

func main() {

}
