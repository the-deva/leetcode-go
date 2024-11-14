package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_minimizedMaximum(t *testing.T) {
	t.Log("Currently testing [2064 : minimizedMaximum]")
	examples := [][]string{
		{
			`6`, `[11,6]`,
			`3`,
		},
		{
			`7`, `[15,10,10]`,
			`5`,
		},
		{
			`1`, `[100000]`,
			`100000`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, minimizedMaximum, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
