/*
1957 -> https://leetcode.com/problems/delete-characters-to-make-fancy-string/description/
1. Update and maintain count should not be greater than 2
2. Build the string maintaining count
*/
package main

import "strings"

func makeFancyString(s string) string {
	var result strings.Builder
	count := 1
	pre := s[0]

	for i := 1; i < len(s); i++ {
		c := s[i]
		if c == pre {
			count++
		} else {
			if count > 2 {
				count = 2
			}
			result.WriteString(strings.Repeat(string(pre), count))
			pre = c
			count = 1
		}
	}
	if count > 2 {
		count = 2
	}
	result.WriteString(strings.Repeat(string(pre), count))
	return result.String()
}

func main() {
	s := "leeetcode"
	makeFancyString(s)
}
