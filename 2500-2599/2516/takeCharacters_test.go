package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_takeCharacters(t *testing.T) {
	t.Log("Currently tesing [2516 : takeCharacters]")
	examples := [][]string{
		{
			`"aabaaaacaabc"`, `2`,
			`8`,
		},
		{
			`"a"`, `1`,
			`-1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, takeCharacters, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
