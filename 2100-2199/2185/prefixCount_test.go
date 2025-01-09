package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_prefixCount(t *testing.T) {
	t.Log("Currently testing [2185 : prefixCount]")
	examples := [][]string{
		{
			`["pay","attention","practice","attend"]`,
			`at`,
		},
		{
			`["leetcode","win","loops","success"]`,
			`code`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, prefixCount, examples, targetCaseNum); err != 0 {
		t.Fatal(err)
	}
}
