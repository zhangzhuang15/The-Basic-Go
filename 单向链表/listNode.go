package main

type ListNode struct {
	value int
	next  *ListNode
}

func New(array []int) *ListNode {
	head := &ListNode{
		value: 0,
		next:  nil,
	}

	cur := head

	for i := 0; i < len(array); i++ {
		node := ListNode{
			value: array[i],
			next:  nil,
		}
		cur.next = &node
		cur = cur.next
	}

	cur = head.next
	head.next = nil

	return cur
}

// 返回 第 k 个节点
func (head *ListNode) GetNode(k int) *ListNode {
	if k <= 0 {
		return nil
	}

	cur := head

	for i := 1; i < k && cur != nil; i++ {
		cur = cur.next
	}

	return cur
}

// 返回倒数 第 k 个节点，
// 双指针算法
func (head *ListNode) GetLastNode(k int) *ListNode {
	if k <= 0 {
		return nil
	}

	left, right := head, head

	for i := 1; i < k && right != nil; i++ {
		right = right.next
	}

	if right == nil {
		return nil
	}

	for right != nil && right.next != nil {
		left = left.next
		right = right.next
	}

	return left
}

// 返回 中间的节点, 中间的节点可能是一个，也可能是两个
// 快慢指针算法
func (head *ListNode) GetMiddleNode() []*ListNode {
	slow, fast := head, head

	for fast != nil && fast.next != nil && fast.next.next != nil {
		// 快指针走2步
		fast = fast.next.next
		// 慢指针走1步
		slow = slow.next
	}

	if fast.next == nil {
		return []*ListNode{slow}
	}

	return []*ListNode{slow, slow.next}
}

// 检测单链表是否有环
func (head *ListNode) HasCircle() bool {
	slow, fast := head, head

	for fast != nil && fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			break
		}
	}
	// 如果没有环，循环肯定结束，而且fast要比slow率先成为 nil，
	// fast如果是nil， slow肯定不是nil，fast 和 slow 肯定不等；
	//
	// 如果有环的话，fast 和 slow 肯定都非 nil，且二者相等

	return fast == slow
}

// 找到单链表中环的起点
func (head *ListNode) GetCircleNode() *ListNode {
	slow, fast := head, head

	for fast != nil && fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}

	if fast != slow {
		return nil
	}

	// 此时 slow == fast, 假设 slow从head来到这个位置前进了k步，那么 fast从head来到这个位置前进了2k步。
	// 那么 k 可能是环上节点的数目，或者是该数目的整数倍。
	// 如果假设环起点前进m步来到这个位置，那么从head来到环起点要前进 k - m 步，
	// 从该位置沿着环继续前进k步的话，就会回到该位置，但继续前进 k - m 步的话，就会来到环的起点位置，
	// 基于这个道理，将 slow 调整到 head， 当下一次 slow 和 fast 相遇的时候，就是环的起点位置。
	slow = head

	for fast != nil && fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next
		if slow == fast {
			break
		}
	}

	if slow == fast {
		return slow
	}

	return nil
}

// 获取环上的节点总数
func (head *ListNode) GetCircleNodeNum() int {
	slow, fast := head, head

	for fast != nil && fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			break
		}
	}

	if fast != slow {
		return 0
	}

	num := 0
	
	for slow != nil && slow.next != nil {
		slow = slow.next
		num += 1
		if slow == fast {
			break
		}
	}

	return num
}

// 获取两个单链表的交叉节点，也就是说，自这个交叉节点开始，后续的节点二者完全相同（节点地址相同）
func GetIntersectionNode(node1 *ListNode, node2 *ListNode) *ListNode {
	i, j := node1, node2

	if i == nil || j == nil {
		return nil
	}

	for i != j {
		if i == nil {
			i = node2
		} else {
			i = i.next
		}

		if j == nil {
			j = node1
		} else {
			j = j.next
		}
	}
	// 如果二者不相交， 那么 i 和 j 会同时等于nil，循环结束，
	// 所以这个循环终将结束

	return i
}
