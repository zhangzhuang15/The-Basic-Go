package main

type SegmentTreeNode struct {
	leftchild  *SegmentTreeNode
	rightchild *SegmentTreeNode
	left       int
	right      int
	value      int
	offset     int
}

// 建立线段树的细节实现
func Build(array []int, left int, right int) *SegmentTreeNode {

	var merge func(left, right int) *SegmentTreeNode
	merge = func(left, right int) *SegmentTreeNode {
		if left > right {
			return nil
		}
		if left == right {
			return &SegmentTreeNode{
				leftchild:  nil,
				rightchild: nil,
				left:       left,
				right:      left,
				value:      array[left],
				offset:     0,
			}
		}

		middle := int((left + right) / 2)

		leftchild := merge(left, middle)
		rightchild := merge(middle+1, right)

		leftvalue, rightvalue := 0, 0

		if leftchild != nil {
			leftvalue = leftchild.value
		}
		if rightchild != nil {
			rightvalue = rightchild.value
		}

		node := &SegmentTreeNode{
			leftchild:  leftchild,
			rightchild: rightchild,
			left:       left,
			right:      right,
			value:      leftvalue + rightvalue,
			offset:     0,
		}

		return node
	}

	root := merge(left, right)

	return root

}

// 创建线段树
func New(array []int) *SegmentTreeNode {
	if len(array) == 0 {
		return nil
	}

	left, right := 0, len(array)-1

	root := Build(array, left, right)

	return root
}

// 索引位于[left, right]区间的元素，每个都加 value
func (root *SegmentTreeNode) Increase(value, left, right int) {

	if left > right {
		return
	}

	// left right 和 root ，在区间上没有交集
	if root == nil || root.left > right || root.right < left {
		return
	}
	// root对应的区间包含于 left right 之间
	if left <= root.left && root.right <= right {
		root.value += (root.right - root.left + 1) * value
		root.offset += value
		return
	}
	// left right 和 root 在区间上只是部分重合，
	// 更新 root 左右子节点
	if root.offset != 0 {
		if root.leftchild != nil {
			root.leftchild.value += (root.leftchild.right - root.leftchild.left + 1) * root.offset
			root.leftchild.offset += root.offset
		}
		if root.rightchild != nil {
			root.rightchild.value += (root.rightchild.right - root.rightchild.left + 1) * root.offset
			root.rightchild.offset += root.offset
		}
		root.offset = 0
	}

	// 递归运算，由左节点增加
	if root.leftchild != nil {
		root.leftchild.Increase(value, left, right)
	}

	// 由右节点增加
	if root.rightchild != nil {
		root.rightchild.Increase(value, left, right)
	}
}

// 查询位于[left, right]区间的元素总和
func (root *SegmentTreeNode) Query(left, right int) int {
	if left > right {
		return 0
	}
	if root == nil || root.left > right || root.right < left {
		return 0
	}

	if left <= root.left && root.right <= right {
		return root.value
	}

	// left right 和 root 在区间上只是部分重合，
	// 更新 root 左右子节点
	if root.offset != 0 {
		if root.leftchild != nil {
			root.leftchild.value += (root.leftchild.right - root.leftchild.left + 1) * root.offset
			root.leftchild.offset += root.offset
		}
		if root.rightchild != nil {
			root.rightchild.value += (root.rightchild.right - root.rightchild.left + 1) * root.offset
			root.rightchild.offset += root.offset
		}
		root.offset = 0
	}

	value := 0

	if root.leftchild != nil {
		value += root.leftchild.Query(left, right)
	}

	if root.rightchild != nil {
		value += root.rightchild.Query(left, right)
	}

	return value
}
