package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_getMaximumXor(t *testing.T) {
	t.Log("Currently testing [1829: getMaximumXor]")
	examples := [][]string{
		{
			`[0,1,1,3]`, `2`,
			`[0,3,2,3]`,
		},
		{
			`[2,3,4,7]`, `3`,
			`[5,2,6,5]`,
		},
		{
			`[0,1,2,2,5,7]`, `3`,
			`[4,3,6,4,6,7]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, getMaximumXor, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
