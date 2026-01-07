package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_maxProduct(t *testing.T) {
	t.Log("Currently testing [maxProduct]")
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithFile(t, maxProduct, "input.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
