package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_maxVowels(t *testing.T) {
	t.Log("Currently testing : [1456 -> maxVowels]")
	examples := [][]string{
		{
			`"abciiidef"`, `3`,
			`3`,
		},
		{
			`"aeiou"`, `2`,
			`2`,
		},
		{
			`"leetcode"`, `3`,
			`2`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxVowels, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
