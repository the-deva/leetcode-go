package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_maximumBeauty(t *testing.T) {
	t.Log("Currently testing [2799: maximumBeauty]")
	examples := [][]string{
		{
			`[4,6,1,2]`, `2`,
			`3`,
		},
		{
			`[1,1,1,1]`, `10`,
			`4`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, maximumBeauty, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
