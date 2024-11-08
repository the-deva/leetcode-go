package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_mergeAlternately(t *testing.T) {
	t.Log("Currently testing : [1768: mergeAlternately]")
	examples := [][]string{
		{
			`"abc"`, `"pqr"`,
			`"apbqcr"`,
		},
		{
			`"ab"`, `"pqrs"`,
			`"apbqrs"`,
		},
		{
			`"abcd"`, `"pq"`,
			`"apbqcd"`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, mergeAlternately, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
