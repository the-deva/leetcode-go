/*
1346 : https://leetcode.com/problems/check-if-n-and-its-double-exist
1. Have a hashMap of all the counts
2. If multiple 0's exist return true else delete the zero
3. See if 2 * k exists in cnts and return true
*/
package main

func checkIfExist(arr []int) bool {
	cnts := map[int]int{}
	for _, v := range arr {
		cnts[v]++
	}

	if cnts[0] > 1 {
		return true
	} // If more than 2 zeros case
	delete(cnts, 0) // remove 0

	for k := range cnts {
		if cnts[2*k] > 0 {
			return true
		}
	}
	return false
}

func main() {

}
