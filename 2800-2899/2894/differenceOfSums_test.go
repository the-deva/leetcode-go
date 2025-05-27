package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_isAcronym(t *testing.T) {
	t.Log("Currently testing [2894 : differenceOfSums]")
	examples := [][]string{
		{
			`10`, `3`,
			`19`,
		},
		{
			`5`, `6`,
			`15`,
		},
		{
			`5`, `1`,
			`-15`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, differenceOfSums, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
