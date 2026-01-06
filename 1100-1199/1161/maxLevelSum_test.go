package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_maxLevelSum(t *testing.T) {
	t.Log("Currently testing [maxLevelSum]")
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithFile(t, maxLevelSum, "input.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
