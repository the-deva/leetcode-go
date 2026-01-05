package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_reversePrefix(t *testing.T) {
	t.Log("Currently testing [3794 -> reversePrefix]")
	examples := [][]string{
		{
			`"abcd"`, `2`,
			`"bacd"`,
		},
		{
			`"xyz"`, `3`,
			`"zyx"`,
		},
		{
			`"hey"`, `1`,
			`"hey"`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, reversePrefix, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
