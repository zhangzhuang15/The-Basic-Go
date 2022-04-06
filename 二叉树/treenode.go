package tree

import (
	"math/rand"
	"sort"
)

type Node struct {
	left *Node
	right *Node
	value int
}

// New
// array 节点数据
// sequence 是否为二叉搜索树
func New(array []int, sequence  bool) *Node {
	if len(array) == 0 {
		return nil
	}

 	var root *Node

	for _, value := range array {
		if root == nil {
			root = &Node{
				value: value,
				left: nil,
				right: nil,
			}
		} else {
			root.Insert(value, sequence)
		}
	}

	return root
}

// Insert 插入新节点
// sequence 是否采用二叉搜索方式插入新节点
func (root *Node)Insert(value int, sequence bool) {
	var prev *Node
	node := root
	// 二叉搜索树插入新节点
	if sequence {
		for node != nil {
			prev = node
			if value < node.value {
				node = node.left
			} else if value > node.value {
				node = node.right
			} else {
				return
			}
		}
		node = &Node{
			value: value,
			left: nil,
			right: nil,
		}
		if value < prev.value {
			prev.left = node
		} else {
			prev.right = node
		}
	} else {   // 随机插入
		for node != nil {
			prev = node
			if rand.Int() % 2 == 0 {
				node = node.left
			} else {
				node = node.right
			}
		}
		node = &Node{
			value: value,
			left: nil,
			right: nil,
		}
		if value < prev.value {
			prev.left = node
		} else {
			prev.right = node
		}
	}
}

// FindAncestorInTwoNodes 寻找最近公共祖先节点
//
// nodeValue1 nodeValue2不一定出现在树当中
func (root *Node)FindAncestorInTwoNodes(nodeValue1, nodeValue2 int) *Node {
	// 如果事先知道 nodeValue1 和 nodeValue2 一定在树中，
	// 应该怎么找出最近公共祖先节点呢？
	// 肯定是从 root 开始遍历，如果当前节点的值刚好等于 nodeValue1 或者 nodeValue2，那么该节点就是要找的节点；
	// 否则的话，再从 root.left 开始遍历找一趟，没找到的话，再从 root.right 开始遍历找一趟。

	// 遗憾的是，nodeValue1 nodeValue2不一定出现在树当中，
	// 按照上面的逻辑寻找的话，找到的节点可能是 nodeValue1的祖先节点，而不是 nodeValue2的祖先节点， 因为 nodeValue2 可能不在树里边。
	// 如果上边说的算法，可以甄别出 nodeValue2不在树中，那么就可以完成求解。
	// 因此，不妨引入两个变量来标记 nodeValue1 nodeValue2 是否在树中，
	// 这就是 findNodeValue1 和 findNodeValue2 的由来。

	// 如果当前节点的值等于 nodeValue1，仍需要访问当前节点的子树，来确定 nodeValue2存不存在，
	// 所以不妨先遍历左右子树，途中就可以知道 nodeValue1 和 nodeValue2 存不存在，就不会有遗漏， 再访问当前节点。

	findNodeValue1, findNodeValue2 := false, false

	var find func(root *Node, val1, val2 int) *Node

	// 假设两个节点都在树中，寻找到可疑祖先节点。
	// 因为这里我们做了假设，找到的节点并不能算是准确的祖先节点，于是称之为 可疑祖先节点。
	find = func(root *Node, val1, val2 int) *Node {
		if root == nil {
			return nil
		}
		leftResult := find(root.left, val1, val2)
		rightResult := find(root.right, val1, val2)

		// 左子树和右子树都找到了可疑最先节点，而 val1 val2在树中只能出现一次，
		// 因此 leftResult 只可能是 val1 和 val2 其中一个的祖先节点， rightResult 同理。
		// 因此 root 被确定为 val1 val2 的 可疑祖先节点
		if leftResult != nil && rightResult != nil {
			return root
		}

		// 剩下的情况有三种：
		// leftResult == nil  rightResult != nil
		// leftResult != nil  rightResult == nil
		// leftResult == nil  rightResult == nil

		// root.value == val1 或者 root.val2成立的话，
		// 表明 leftResult == nil  rightResult == nil
		// 因为 如果 leftResult != nil ， 那么就说明 val1 和 val2 出现在 root 的左子树中，
		// 而树中的值不重复，那么 root 的值肯定不是 val1 和 val2.
		// rightResult != nil 同理。
		// 这时候就要做标记，而 root 本身就可以作为可疑祖先节点返回。
		// root子树的情况，已经在上边的两个递归中标记过了。
		if root.value == val1 {
			findNodeValue1 = true
			return root
		}
		if root.value == val2 {
			findNodeValue2 = true
			return root
		}

		// 到这里只剩下两种情况：
		// leftResult == nil  rightResult != nil
		// leftResult != nil  rightResult == nil
		// 那么 非 nil 的节点肯定就可疑祖先节点了
		if leftResult != nil {
			return leftResult
		}
		return rightResult
	}

	node := find(root, nodeValue1, nodeValue2)
	// node只是可疑祖先节点，
	// 如果  findNodeValue1 findNodeValue2 都是 true，
	// 就表明两个节点都在树中，
	// 那么find方法中找到的节点肯定就是祖先节点了。
	// 否则，表明至少有一个节点不在树中，那么祖先节点就是 nil 了
	if !findNodeValue1 || !findNodeValue2 {
		return nil
	}
	return node
}

// FindAncestorInSliceNodes 寻找array中所有元素表示的节点的最近公共祖先节点
func (root *Node)FindAncestorInSliceNodes(array []int) *Node {
    if len(array) == 0 {
    	return nil
	}
	// 记录 array 中的元素是否位于树中
	hashMap := make(map[int]bool)

	// O(1)时间内判断 root.value 是否在 array 中
	hashSet := make(map[int]bool)
	for _, value := range array {
		hashSet[value] = true
	}

	var find func(root *Node) *Node

	find = func(root *Node) *Node {
		if root == nil {
			return nil
		}

		leftResult := find(root.left)
		rightResult := find(root.right)

		if leftResult != nil && rightResult != nil {
			return root
		}

		// root.value in array
		// 两个值的情形下，是用 root.value分别去比较，
		// 多个值的情形下，直接拓展为集合判断。
		if _, ok := hashSet[root.value]; ok {
			hashMap[root.value] = true
			return root
		}

		if leftResult != nil {
			return leftResult
		}
		return rightResult
	}

	node := find(root)

	// 因为 array 中可能存在重复的值，因此使用 len(hashSet)
	// array中的元素全部在树中
	if len(hashMap) == len(hashSet) {
		return node
	}
	return nil
}

// FindAncestorInTwoNodesBetter 利用二叉搜索树的顺序特性优化查找最近公共祖先节点
func (root *Node)FindAncestorInTwoNodesBetter(nodeValue1, nodeValue2 int) *Node {
	if nodeValue1 > nodeValue2 {
		nodeValue1, nodeValue2 = nodeValue2, nodeValue1
	}
	findNodeValue1, findNodeValue2 := false, false
	var find func(root *Node) *Node

	find = func(root *Node) *Node {
		if root == nil {
			return nil
		}

		if nodeValue1 > root.value {
			return find(root.right)
		} else if nodeValue2 < root.value {
			return find(root.left)
		} else {
			leftResult := find(root.left)
			rightResult := find(root.right)

			if leftResult != nil && rightResult != nil {
				return root
			}

			if root.value == nodeValue1 {
				findNodeValue1 = true
				return root
			}

			if root.value == nodeValue2 {
				findNodeValue2 = true
				return root
			}

			if leftResult != nil {
				return leftResult
			}
			return rightResult
		}
	}

	node := find(root)

	if !findNodeValue1 || !findNodeValue2 {
		return nil
	}
	return node
}

func (root *Node)FindAncestorInSliceNodesBetter(array []int) *Node {
	if len(array) == 0 {
		return nil
	}

	sort.Ints(array)
	minValue, maxValue := array[0], array[len(array)-1]

	// 记录 array 中的元素是否位于树中
	hashMap := make(map[int]bool)

	// O(1)时间内判断 root.value 是否在 array 中
	hashSet := make(map[int]bool)
	for _, value := range array {
		hashSet[value] = true
	}

	var find func(root *Node) *Node

	find = func(root *Node) *Node {
		if root == nil {
			return nil
		}

		if minValue > root.value {
			return find(root.right)
		} else if maxValue < root.value {
			return find(root.left)
		} else {
			leftResult := find(root.left)
			rightResult := find(root.right)

			if leftResult != nil && rightResult != nil {
				return root
			}

			// root.value in array
			// 两个值的情形下，是用 root.value分别去比较，
			// 多个值的情形下，直接拓展为集合判断。
			if _, ok := hashSet[root.value]; ok {
				hashMap[root.value] = true
				return root
			}

			if leftResult != nil {
				return leftResult
			}
			return rightResult
		}
	}

	node := find(root)

	// 因为 array 中可能存在重复的值，因此使用 len(hashSet)
	// array中的元素全部在树中
	if len(hashMap) == len(hashSet) {
		return node
	}
	return nil
}
