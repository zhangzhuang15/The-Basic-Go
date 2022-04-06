package main

import "testing"

func TestIncrease(t *testing.T) {

	array := []int{2, 4, 1, 8, 9}

	root := New(array)

	root.Increase(2, 0, 2)

	if root.Query(0, 2) != 13 {
		t.FailNow()
	}
}

func TestQuery(t *testing.T) {

	array := []int{2, 4, 1, 8, 9}

	root := New(array)

	if root.Query(0, 2) == 7 &&
		root.Query(2, 4) == 18 &&
		root.Query(1, 2) == 5 &&
		root.Query(-1, 0) == 2 &&
		root.Query(4, 5) == 9 &&
		root.Query(5, 1) == 0 &&
		root.Query(-3, -2) == 0 &&
		root.Query(5, 10) == 0 {
		return
	}
	t.FailNow()
}
