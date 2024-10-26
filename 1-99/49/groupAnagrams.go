package main  

import ("github.com/the-chowdary/leetcode-go/utils")

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
    hashMap := make(map[string][]string)

	for _, word := range strs {
		sortedWord := utils.sortString(word)
		hashMap[sortedWord] = append(hashMap[sortedWord], word)
	}

	result := make([][]string, 0, len(hashMap))
	for _, wordsList := range hashMap {
		result = append(result, wordsList)
	}
	return result
}