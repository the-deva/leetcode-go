package main

// https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/description/

func MaxUniqueSplit(s string) (result int) {
   	has := map[string]bool{}
	var dfs func(string)
	dfs = func(s string) {
		if s == "" {
			if len(s) > result {
				result = len(s)
			}
			return
		}
		for i := 1; i < len(s); i++ {
			if t := s[:i]; !has[t] {
				has[t] = true
				dfs(s[i:])
				delete(has, t)
			}
		}
	}
	dfs(s)
	return
}

