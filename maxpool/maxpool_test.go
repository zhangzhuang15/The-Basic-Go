package maxpool

import "testing"

func TestMaxPoolWithSlideOne(t *testing.T) {
	result := MaxPool([]int{2, 4, 1, 5, 8, 7, 3}, 4, 1)
	if len(result) != 4 {
		t.FailNow()
	}
	if result[0] != 5 ||
		result[1] != 8 ||
		result[2] != 8 ||
		result[3] != 8 {
		t.FailNow()
	}
}

func TestMaxPoolWithSlideTwo(t *testing.T) {
	result := MaxPool([]int{ 2, 8, 4, 1, 5, 7, 3}, 3, 2)
	if len(result) != 3 {
		t.FailNow()
	}
	if result[0] != 8 ||
		result[1] != 5 ||
		result[2] != 7 {
		t.FailNow()
	}
}

func TestMaxPoolWithSlideThree(t *testing.T) {
	result := MaxPool([]int{2, 8, 4, 1, 6, 7, 5, 3, 9}, 3, 3 )
	if len(result) != 3 {
		t.FailNow()
	}
	if result[0] != 8 ||
		result[1] != 7 ||
		result[2] != 9 {
		t.FailNow()
	}
}
