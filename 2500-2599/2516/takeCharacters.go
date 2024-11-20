/*
2516 : https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/description
1. Sliding window
*/
package main

func takeCharacters(s string, k int) int {
	cnt := [3]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	if cnt[0] < k || cnt[1] < k || cnt[2] < k {
		return -1
	}

	mx, left := 0, 0
	for right, c := range s {
		c -= 'a'
		cnt[c]--
		for cnt[c] < k {
			cnt[s[left]-'a']++
			left++
		}
		mx = max(mx, right-left+1)
	}
	return len(s) - mx
}

func main() {

}
