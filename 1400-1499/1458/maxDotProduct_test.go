package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_maxDotProduct(t *testing.T) {
	t.Log("Currently testing : [1458 -> maxDotProduct]")
	examples := [][]string{
		{
			`[2,1,-2,5]`, `[3,0,-6]`,
			`18`,
		},
		{
			`[3,-2]`, `[2,-6,7]`,
			`21`,
		},
		{
			`[-1,-1]`, `[1,1]`,
			`-1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, maxDotProduct, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
