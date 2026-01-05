/*
3794 - https://leetcode.com/problems/reverse-string-prefix/
*/
package main

func reversePrefix(s string, k int) string {
	runes := []rune(s)
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {

}
