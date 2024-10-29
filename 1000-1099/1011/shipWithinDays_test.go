package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_shipWithinDays(t *testing.T) {
	t.Log("Currently testing [1011 : shipWithinDays]")
	examples := [][]string{
		{
			`[1,2,3,4,5,6,7,8,9,10]`, `5`,
			`15`,
		},
		{
			`[3,2,2,4,1,4]`, `3`,
			`6`,
		},
		{
			`[1,2,3,1,1]`, `4`,
			`3`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, shipWithinDays, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
