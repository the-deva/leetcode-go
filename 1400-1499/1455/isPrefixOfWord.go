/*
1455 : https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/description/
1. Split the sentence and verify prefix
*/
package main

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	for i, s := range strings.Split(sentence, " ") {
		if strings.HasPrefix(s, searchWord) {
			return i + 1
		}
	}
	return -1
}

func main() {

}
