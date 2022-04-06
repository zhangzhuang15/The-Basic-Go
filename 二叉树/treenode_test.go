package tree

import "testing"

func TestNode_FindAncestorInTwoNodes(t *testing.T) {
	root := New([]int{1, 3, 2, 4, 5, 7, 6}, true)

	node := root.FindAncestorInTwoNodes(7, 6)
	if node == nil || node.value != 7 {
		t.FailNow()
	}

	node = root.FindAncestorInTwoNodes(5, 6)
	if node == nil || node.value != 5 {
		t.FailNow()
	}

	node = root.FindAncestorInTwoNodes(2, 5)
	if node == nil || node.value != 3 {
		t.FailNow()
	}

	node = root.FindAncestorInTwoNodes(2, 0)
	if node != nil {
		t.FailNow()
	}

	node = root.FindAncestorInTwoNodes(-2, 18)
	if node != nil {
		t.FailNow()
	}
}

func TestNode_FindAncestorInSliceNodes(t *testing.T) {
	root := New([]int{3, 2, 5, 4, 9, 1, 6}, true)

	node := root.FindAncestorInSliceNodes([]int{1, 4, 5})
	if node == nil || node.value != 3 {
		t.FailNow()
	}

	node = root.FindAncestorInSliceNodes([]int{2 , 1, 3})
	if node == nil || node.value !=3 {
		t.FailNow()
	}

	node = root.FindAncestorInSliceNodes([]int{1, 6, 5})
	if node == nil || node.value !=3 {
		t.FailNow()
	}

	node = root.FindAncestorInSliceNodes([]int{1, 6, 10})
	if node != nil {
		t.FailNow()
	}

	node = root.FindAncestorInSliceNodes([]int{0, 8, -2, -2, 7, 0})
	if node != nil {
		t.FailNow()
	}

	node = root.FindAncestorInSliceNodesBetter([]int{})
	if node != nil {
		t.FailNow()
	}
}

func TestNode_FindAncestorInTwoNodesBetter(t *testing.T) {
	root := New([]int{1, 3, 2, 4, 5, 7, 6}, true)

	node := root.FindAncestorInTwoNodesBetter(7, 6)
	if node == nil || node.value != 7 {
		t.FailNow()
	}

	node = root.FindAncestorInTwoNodesBetter(5, 6)
	if node == nil || node.value != 5 {
		t.FailNow()
	}

	node = root.FindAncestorInTwoNodesBetter(2, 5)
	if node == nil || node.value != 3 {
		t.FailNow()
	}

	node = root.FindAncestorInTwoNodesBetter(2, 0)
	if node != nil {
		t.FailNow()
	}

	node = root.FindAncestorInTwoNodesBetter(-2, 18)
	if node != nil {
		t.FailNow()
	}
}

func TestNode_FindAncestorInSliceNodesBetter(t *testing.T) {
	root := New([]int{3, 2, 5, 4, 9, 1, 6}, true)

	node := root.FindAncestorInSliceNodesBetter([]int{1, 4, 5})
	if node == nil || node.value != 3 {
		t.FailNow()
	}

	node = root.FindAncestorInSliceNodesBetter([]int{2 , 1, 3})
	if node == nil || node.value !=3 {
		t.FailNow()
	}

	node = root.FindAncestorInSliceNodesBetter([]int{1, 6, 5})
	if node == nil || node.value !=3 {
		t.FailNow()
	}

	node = root.FindAncestorInSliceNodesBetter([]int{1, 6, 10})
	if node != nil {
		t.FailNow()
	}

	node = root.FindAncestorInSliceNodesBetter([]int{0, 8, -2, -2, 7, 0})
	if node != nil {
		t.FailNow()
	}

	node = root.FindAncestorInSliceNodesBetter([]int{})
	if node != nil {
		t.FailNow()
	}
}
