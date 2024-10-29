package utils

import "testing"

func TestSortUtils(t *testing.T) {
	got := SortString("cab")
	want := "abc"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestMax(t *testing.T) {
	got := Max([]int{1, 2, 3, 4, 5})
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3, 4, 5})
	want := 15

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
