/*
3289 : https://leetcode.com/problems/the-two-sneaky-numbers-of-digitville/description/
1. Calculate the count of each number in the input array.
2. Iterate through the input array and check the count of each number.
3. If the count of a number is 2 or more, add its index to the result array.
*/
package main

import "sort"

func getSneakyNumbers(nums []int) []int {
	numsCount := make(map[int]int)
	for _, num := range nums {
		numsCount[num]++
	}

	result := make([]int, 0)
	for i, num := range numsCount {
		if num >= 2 {
			result = append(result, i)
		}
	}
	sort.Ints(result)
	return result
}

func main() {

}
