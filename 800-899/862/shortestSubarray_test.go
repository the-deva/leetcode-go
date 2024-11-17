package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_shortestSubarray(t *testing.T) {
	t.Log("Current test is [862 : shortestSubarray]")
	examples := [][]string{
		{
			`[1]`, `1`,
			`1`,
		},
		{
			`[1,2]`, `4`,
			`-1`,
		},
		{
			`[2,-1,2]`, `3`,
			`3`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, shortestSubarray, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
