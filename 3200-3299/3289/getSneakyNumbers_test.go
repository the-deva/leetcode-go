package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_getSneakyNumbers(t *testing.T) {
	t.Log("Currently testing [3289 : getSneakyNumbers]")
	examples := [][]string{
		{
			`[0,1,1,0]`,
			`[0,1]`,
		},
		{
			`[0,3,2,1,3,2]`,
			`[2,3]`,
		},
		{
			`[7,1,5,4,3,4,6,0,9,5,8,2]`,
			`[4,5]`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, getSneakyNumbers, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
