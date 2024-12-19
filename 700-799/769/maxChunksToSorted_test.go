package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_maxChunksToSorted(t *testing.T) {
	t.Log("Currently testing [769 : maxChunksToSorted]")
	examples := [][]string{
		{
			`[4,3,2,1,0]`,
			`1`,
		},
		{
			`[1,0,2,3,4]`,
			`4`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxChunksToSorted, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
