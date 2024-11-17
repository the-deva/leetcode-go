/*
862 : https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/
1. Calculate prefix sum
2. If valid subarray found then update result
3. Maintain the stack strictly increasing.
4. Return result
*/
package main

import "container/list"

func shortestSubarray(nums []int, k int) int {
	N := len(nums)
	result := N + 1
	stack := list.New()
	stack.PushBack([]int{0, -1}) // pre = 0 and index = -1
	pre := 0

	for i := 0; i < N; i++ {
		pre += nums[i]

		// valid subarray if sum >= k
		for stack.Len() > 0 && stack.Front().Value.([]int)[0] <= pre-k {
			j := stack.Front().Value.([]int)[1]
			stack.Remove(stack.Front())
			if i-j < result {
				result = i - j
			}
		}

		// Maintain strictly increasing stack
		for stack.Len() > 0 && stack.Back().Value.([]int)[0] >= pre {
			stack.Remove(stack.Back())
		}

		stack.PushBack([]int{pre, i})
	}

	if result < N+1 {
		return result
	}
	return -1
}

func main() {

}
