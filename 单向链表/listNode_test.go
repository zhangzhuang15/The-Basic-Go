package main

import "testing"

func TestGetNode(t *testing.T) {
	head := New([]int{1, 4, 3, 5, 6})

	node := head.GetNode(3)
	if node.value != 3 {
		t.FailNow()
	}

	node = head.GetNode(0)
	if node != nil {
		t.FailNow()
	}

	node = head.GetNode(-3)
	if node != nil {
		t.FailNow()
	}
}

func TestGetLastNode(t *testing.T) {
	head := New([]int{1, 4, 3, 5, 6})

	node := head.GetLastNode(2)
	if node.value != 5 {
		t.Fail()
	}

	node = head.GetLastNode(0)
	if node != nil {
		t.Fail()
	}

	node = head.GetLastNode(-4)
	if node != nil {
		t.Fail()
	}

	node = head.GetLastNode(6)
	if node != nil {
		t.Fail()
	}
}

func TestGetMiddleNode_Odd(t *testing.T) {
	head := New([]int{1, 4, 3, 5, 6})

	nodes := head.GetMiddleNode()
	if len(nodes) == 0 || nodes[0].value != 3 {
		t.Fail()
	}
}

func TestGetMiddleNode_Even(t *testing.T) {
	head := New([]int{1, 4, 3, 5, 6, 8})

	nodes := head.GetMiddleNode()
	if len(nodes) != 2 || nodes[0].value != 3 || nodes[1].value != 5 {
		t.Fail()
	}
}

func TestHasCircle(t *testing.T) {
	head := New([]int{1, 4, 3, 5, 6})

	if head.HasCircle() {
		t.Fail()
	}

	cur := head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = head

	if head.HasCircle() == false {
		t.Fail()
	}
}

func TestGetCircleNode(t *testing.T) {
	head := New([]int{1, 4, 3, 5, 6})

	node := head.GetCircleNode()
	if node != nil {
		t.Fail()
	}

	cur := head
	var p *ListNode
	for cur.next != nil {
		if cur.value == 3 {
			p = cur
		}
		cur = cur.next
	}
	cur.next = p

	node = head.GetCircleNode()
	if node == nil || node.value != 3 {
		t.Fail()
	}
}

func TestIntersectionNode(t *testing.T) {
	head := New([]int{1, 4, 3, 5, 6})

	cur := head.GetNode(3)

	newHead := New([]int{8, 2})
	tail := newHead.GetLastNode(1)
	tail.next = cur

	node := GetIntersectionNode(head, newHead)
	if node == nil || node != cur {
		t.Fail()
	}
}
