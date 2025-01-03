package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_waysToSplitArray(t *testing.T) {
	t.Log("Currently testing [2270 : waysToSplitArray]")
	examples := [][]string{
		{
			`[10,4,-8,7]`,
			`2`,
		},
		{
			`[2,3,1,0]`,
			`2`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, waysToSplitArray, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
