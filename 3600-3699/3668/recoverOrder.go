/*
3668 - https://leetcode.com/problems/restore-finishing-order/description/
*/
package main

func recoverOrder(order, friends []int) []int {
	N := len(order)
	isFriend := make([]bool, N+1)
	for _, x := range friends {
		isFriend[x] = true
	}

	result := make([]int, 0, len(friends))
	for _, x := range order {
		if isFriend[x] {
			result = append(result, x)
		}
	}
	return result
}

func main() {

}
