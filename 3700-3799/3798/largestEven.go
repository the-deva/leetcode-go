/*
3798 - https://leetcode.com/problems/largest-even-number/
1. Trim all trailing '1's from the input binary string to form the largest even number.
*/
package main

import "strings"

func largestEven(s string) string {
	return strings.TrimRight(s, "1")
}

func main() {

}
