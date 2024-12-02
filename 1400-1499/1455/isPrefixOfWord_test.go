package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_isPrefixOfWord(t *testing.T) {
	t.Log("Currently testing [1455 : isPrefixOfWord]")
	examples := [][]string{
		{
			`"i love eating burger"`, `"burg"`,
			`4`,
		},
		{
			`"this problem is an easy problem"`, `"pro"`,
			`2`,
		},
		{
			`"i am tired"`, `"you"`,
			`-1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, isPrefixOfWord, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
