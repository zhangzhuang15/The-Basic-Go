package scattertable

import "testing"

func TestMaxIn_Odd(t *testing.T) {
	arr := []int { 3, 2, 1, 6, 8, 9, 21}

	result := MaxIn(arr, 1, 5)

	if result != 8 {
		t.Fail()
	}
}

func TestMaxIn_Even(t *testing.T) {
	arr := []int { 2, 7, 9, 0, 3, 8, 1, 11}

	result := MaxIn(arr, 2, 7)

	if result != 9 {
		t.Fail()
	}
}
