package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_findChampion(t *testing.T) {
	t.Log("Currently testing [2924 : findChampion]")
	examples := [][]string{
		{
			`3`, `[[0,1],[1,2]]`,
			`0`,
		},
		{
			`4`, `[[0,2],[1,3],[1,2]]`,
			`-1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, findChampion, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
