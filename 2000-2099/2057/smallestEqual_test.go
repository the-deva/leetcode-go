package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_smallestEqual(t *testing.T) {
	t.Log("Currently testing [2057: smallestEqual]")
	examples := [][]string{
		{
			`[0,1,2]`,
			`0`,
		},
		{
			`[4,3,2,1]`,
			`2`,
		},
		{
			`[1,2,3,4,5,6,7,8,9,0]`,
			`-1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, smallestEqual, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
