package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_recoverOrder(t *testing.T) {
	t.Log("Currently testing [3512 -> recoverOrder]")
	examples := [][]string{
		{
			`[3,1,2,5,4]`, `[1,3,4]`,
			`[3,1,4]`,
		},
		{
			`[1,4,5,3,2]`, `[2,5]`,
			`[5,2]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, recoverOrder, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
