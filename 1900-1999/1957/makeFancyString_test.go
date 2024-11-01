package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_makeFancyString(t *testing.T) {
	t.Log("Currently testing [1957 : makeStringFancy]")
	examples := [][]string{
		{
			`"leeetcode"`,
			`"leetcode"`,
		},
		{
			`"aaabaaaa"`,
			`"aabaa"`,
		},
		{
			`"aab"`,
			`"aab"`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, makeFancyString, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
