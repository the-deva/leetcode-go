package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_longestSquareStreak(t *testing.T) {
	t.Log("Currently testing [2501 -> longestSquareStreak]")
	examples := [][]string{
		{
			`[4,3,6,16,8,2]`,
			`3`,
		},
		{
			`[2,3,5,6,7]`,
			`-1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, longestSquareStreak, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
