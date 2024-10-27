package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_robotSim(t *testing.T) {
	t.Log("Currently testing [874 -> robotSim]")
	examples := [][]string{
		{
			`[4,-1,3]`, `[]`,
			`25`,
		},
		{
			`[4,-1,4,-2,4]`, `[[2,4]]`,
			`65`,
		},
		{
			`[6,-1,-1,6]`, `[[0,0]]`,
			`36`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, robotSim, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
