/*
3330 : https://leetcode.com/problems/find-the-original-typed-string-i/description/
1. There will be one string (result = 1)
2. Iterate the word if word[i - 1] == word[i] then result++
*/
package main

func possibleStringCount(word string) int {
	result := 1
	for i := 1; i < len(word); i++ {
		if word[i-1] == word[i] {
			result++
		}
	}
	return result
}

func main() {

}
