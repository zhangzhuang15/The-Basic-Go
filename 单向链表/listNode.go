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

// 返回第 k 个节点
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

// 返回倒数第 k 个节点
//
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

// 返回中间的节点, 中间的节点可能是一个，也可能是两个
//
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

// 单项链表就地排序，小 -> 大
// O(nlogn)
func Sort(start *ListNode, up bool) *ListNode {

	length := 0

	// 获取链表长度
	for cur := start; cur != nil; length, cur = length+1, cur.next {}

	// 哨兵节点
	t := ListNode{0, start}

	// prev 是有序链表的末尾节点
	// cur  是待排序的第一个节点
	// next 是 cur 排序后要继续的位置
	prev, cur := &t, start
	var next *ListNode = nil
    
	subLen := 1

	// subLen 迭代，省略归并算法中的分组递归
	// subLen 一定是 ≤，避免长度为 2 的整数倍时，漏掉一次排序
	for ; subLen <= length/2; subLen *= 2 {
		for cur != nil {
            // 思路是，从 cur 开始数 2*subLen 个节点，分成两个子链表，两个子链表按顺序合并

			// first 第一个链表的头节点
			// second 第二个链表的头节点
			first := cur
			var second *ListNode = nil

			// 确定第二个链表头节点

			for i := 1; i < subLen && cur != nil && cur.next != nil; i += 1 {
				cur = cur.next
			}
			// 确定第二个链表头节点
			if cur != nil {
				second = cur.next
				// 断开第二个链表头节点
				// 必须要断开，这样在遍历第一个链表时，当节点为 nil ，就可以判断第一个链表结束
				cur.next = nil
			} else {
				second = nil
			}

			// 确定 next
			node := second
			for i := 1; i < subLen && node != nil && node.next != nil; i += 1 {
				node = node.next
			}
			if node != nil {
				next = node.next
				node.next = nil
			} else {
				next = nil
			}

			// first 和 second  合并
			for first != nil && second != nil {
				// 升序
				// 注意 prev 的赋值位置要放在后边，否则会有问题
				if up {
					smaller := first
					if first.value > second.value {
						smaller = second
						second = second.next
					} else {
						first = first.next
					}
					prev.next = smaller
					prev = prev.next
					
					// 降序
				} else {
					bigger := first
					if first.value < second.value {
						bigger = second
						second = second.next
					} else {
						first = first.next
					}
					prev.next = bigger
					prev = prev.next
				}
			}

			// 选中非nil节点
			notNilOne := first
			if second != nil {
				notNilOne = second
			}

			prev.next = notNilOne

			for prev.next != nil {
				prev = prev.next
			}
			cur = next
		}
		prev, cur = &t, t.next
	}

	return t.next
}
