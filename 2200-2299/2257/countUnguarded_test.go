package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_countUnguarded(t *testing.T) {
	t.Log("Currently testing [2257 : countUnguarded]")
	examples := [][]string{
		{
			`4`, `6`, `[[0,0],[1,1],[2,3]]`, `[[0,1],[2,2],[1,4]]`,
			`7`,
		},
		{
			`3`, `3`, `[[1,1]]`, `[[0,1],[1,0],[2,1],[1,2]]`,
			`4`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, countUnguarded, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
