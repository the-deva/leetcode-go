package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_possibleStringCount(t *testing.T) {
	t.Log("Currently testing [3330 -> possibleStringCount]")
	examples := [][]string{
		{
			`"abbcccc"`,
			`5`,
		},
		{
			`"abcd"`,
			`1`,
		},
		{
			`"aaaa"`,
			`4`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, possibleStringCount, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
