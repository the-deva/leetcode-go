package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	t.Log("Currently testing [485 : findMaxConsecutiveOnes]")
	examples := [][]string{
		{
			`[1,1,0,1,1,1]`,
			`3`,
		},
		{
			`[1,0,1,1,0,1]`,
			`2`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, findMaxConsecutiveOnes, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
