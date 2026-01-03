package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_transformArray(t *testing.T) {
	t.Log("Currently testing [3400 : transformArray]")
	examples := [][]string{
		{
			`[4,3,2,1]`,
			`[0,0,1,1]`,
		},
		{
			`[1,5,1,4,2]`,
			`[0,0,1,1,1]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, transformArray, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
