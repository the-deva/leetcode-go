package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_largestEven(t *testing.T) {
	t.Log("Currently testing [3798 -> largestEven]")
	examples := [][]string{
		{
			`"1112"`,
			`"1112"`,
		},
		{
			`"221"`,
			`"22"`,
		},
		{
			`"1"`,
			`""`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, largestEven, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
