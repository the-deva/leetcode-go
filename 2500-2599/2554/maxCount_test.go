package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_maxCount(t *testing.T) {
	t.Log("Currently testing [2554 : maxCount]")
	examples := [][]string{
		{
			`[1,6,5]`, `5`, `6`,
			`2`,
		},
		{
			`[1,2,3,4,5,6,7]`, `8`, `1`,
			`0`,
		},
		{
			`[11]`, `7`, `50`,
			`7`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxCount, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
