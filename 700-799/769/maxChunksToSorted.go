/*
769 : https://leetcode.com/problems/max-chunks-to-make-sorted/description
*/
package main

func maxChunksToSorted(arr []int) (result int) {
	current := 0
	for i, num := range arr {
		if current < num {
			current = num
		}
		if current == i {
			result++
		}
	}
	return
}

func main() {

}
