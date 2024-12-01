package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_checkIfExist(t *testing.T) {
	t.Log("Currently testing [1346 : checkIfExist]")
	examples := [][]string{
		{
			`[10,2,5,3]`,
			`true`,
		},
		{
			`[3,1,7,11]`,
			`false`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, checkIfExist, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
