/*
2554 : https://leetcode.com/problems/maximum-number-of-integers-to-choose-from-a-range-i/description
1. Have a map with all banned numbers
2. Iterate from 1 to n and if i exists in banned list skip
3. else subtract maxSum -= i and result++ and return result
*/
package main

func maxCount(banned []int, n, maxSum int) (result int) {
	has := map[int]bool{}
	for _, v := range banned {
		has[v] = true
	}

	for i := 1; i <= n && i <= maxSum; i++ {
		if !has[i] {
			maxSum -= i
			result++
		}
	}
	return
}

func main() {

}
