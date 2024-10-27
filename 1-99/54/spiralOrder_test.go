package main

import (
	"leetcode-go/testutil"
	"testing"
)

func Test_SpiralOrder(t *testing.T) {
	examples := [][]string{
		{
			`[[1,2,3],[4,5,6],[7,8,9]]`,
			`[1,2,3,6,9,8,7,4,5]`,
		},
		{
			`[[1,2,3,4],[5,6,7,8],[9,10,11,12]]`,
			`[1,2,3,4,8,12,11,10,9,5,6,7]`,
		},
	}
	t.Run("Currently testing [54 -> spiralMatrix modifying source matrix]", func(t *testing.T) {
		targetCaseNum := 0
		if err := testutil.RunLeetCodeFuncWithExamples(t, spiralOrder, examples, targetCaseNum); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Currently testing [54 -> spiralMatrix without modifying source matrix]", func(t *testing.T) {
		targetCaseNum := 0
		if err := testutil.RunLeetCodeFuncWithExamples(t, sprialOrder2, examples, targetCaseNum); err != nil {
			t.Fatal(err)
		}
	})
}
