package main

import "fmt"

// b树：
// 1. 有一个最小度概念，假设最小度 N = 3, 意味着：
//     i 除去根节点，所有节点最少拥有 N 个 子节点，最多拥有 2 * N 个节点，最少拥有 N - 1 个key，最多拥有 2 * N - 1 个key；
//
// 2. 新的key必须插入到叶节点；
// 3. 每个节点同时存储 key 和 值data；（b+树中，data只存储到叶节点，内部节点只存储 key）
//

// b树的实现参考《算法导论》（第三版）

type Node struct {
	nKey     uint
	nChild   uint
	keys     []rune
	data     []interface{}
	children []*Node
	isLeaf   bool
}

type Tree struct {
	minDegree uint
	maxDegree uint
	root      *Node
}

type Result struct {
	node       *Node
	indexOfKey uint
}

func NewNode(minDegree uint) *Node {
	return &Node{
		nKey:     0,
		nChild:   0,
		keys:     make([]rune, 2*minDegree-1),
		data:     make([]interface{}, 2*minDegree-1),
		children: make([]*Node, 2*minDegree),
		isLeaf:   true,
	}
}

func NewTree(minDegree uint) *Tree {
	tree := &Tree{
		minDegree: minDegree,
		maxDegree: 2 * minDegree,
		root:      nil,
	}

	tree.root = NewNode(minDegree)

	return tree
}

// 节点的key是否满了
func isFull(node *Node, minDegree uint) bool {
	if node.nKey >= 2*minDegree-1 {
		return true
	}
	return false
}

// 二分搜索数组
func search(array []rune, key rune) int {
	// 采取 [) 的搜索思路
	start, end := 0, len(array)

	// [ 右移，排除肯定不符合要求的选择，
	// ) 左移，）的位置是 可能符合要求的选择
	for start < end {
		middle := int((start + end) / 2)
		if array[middle] < key {
			start = middle + 1
		} else {
			end = middle
		}
	}

	if array[start] == key {
		return start
	}

	return -1
}

// 节点的key到达下限
func isTooShort(node *Node, minDegree uint) bool {
	if node.nKey <= minDegree-1 {
		return true
	}
	return false
}

func traverse(node *Node, deal func(key rune)) {

	queue := make([]*Node, 0)
	queue = append(queue, node)

	height := 0

	for len(queue) > 0 {
        println("height: ", height)
		for n := len(queue);n > 0; n -= 1 {
			r := queue[0]
			copy(queue[0:], queue[1:])
			queue = queue[0: len(queue) - 1]

			if r != nil {
				for i := 0; i < int(r.nKey); i += 1 {
					deal(r.keys[i])
				}
				print("\t")
				queue = append(queue, r.children[0: r.nKey + 1]...)

			}	
		}
		fmt.Printf("\nlevel ---------------\n")
		height += 1
	}
}

func (root *Node) search(key rune) *Result {
	i := uint(0)

	for i < uint(root.nKey) && root.keys[i] < key {
		i += 1
	}

	if i < uint(root.nKey) && root.keys[i] == key {
		return &Result{
			node:       root,
			indexOfKey: i,
		}
	} else if root.isLeaf {
		return nil
	}

	r := root.children[i]
	return r.search(key)
}

// 节点内，是否包含key
func (root *Node) has(key rune) bool {
	for i := 0; i < int(root.nKey); i += 1 {
		if root.keys[i] == key {
			return true
		}
	}
	return false
}

// 返回 root 的 包含 key 的子节点的索引
func (root *Node) childIndexHasKey(key rune) int {
	if root.has(key) {
		return -1
	}

	for i := 0; i < int(root.nChild); i += 1 {
		child := root.children[i]

		if child.has(key) {
			return i
		}
	}

	return -1
}

// root.children[i] 是满的， root不是满的，将 root.children[i] 拆分
func (root *Node) split(i uint, minDegree uint) {
	child := root.children[i]
	middle := uint(child.nKey / 2)

	node := NewNode(minDegree)

	// 将 child.keys[middle] 右侧的数据植入 node 中
	for j := uint(0); j < 2*minDegree-1-1-middle; j += 1 {
		node.keys[j] = child.keys[j+middle+1]
		node.children[j] = child.children[j+middle+1]
		node.nChild += 1
		node.nKey += 1
		child.nChild -= 1
		child.nKey -= 1
	}
	node.children[2 * minDegree - 1 - 1 - middle] = child.children[2 * minDegree - 1]

	// 将 child.keys[middle]提升到 root 中
	for j := root.nKey; j > i; j -= 1 {
		root.keys[j] = root.keys[j - 1]
		root.data[j] = root.data[j - 1]
		// 注意 child.keys[j] 右侧有个子节点指针，child.children[j + 1]
		// child.children[j] 是 child.keys[j]左侧的子节点指针
		root.children[j+1] = root.children[j]
	}

	root.keys[i] = child.keys[middle]
	root.data[i] = child.data[middle]
	root.nKey += 1
	child.nKey -= 1

	root.children[i+1] = node
	root.nChild += 1
}

// root 节点key不是满的，才能执行这个方法
func (root *Node) insert(key rune, value interface{}, minDegree uint) bool {
	i := uint(0)

	for i < root.nKey && root.keys[i] < key {
		i += 1
	}

	if i < root.nKey && root.keys[i] == key {
		root.data[i] = value
		return true
	}

	// 叶子节点直接插
	if root.isLeaf {
		for j := root.nKey; j > i; j -= 1 {
			root.keys[j] = root.keys[j - 1]
			root.data[j] = root.data[j - 1]
		}
		root.keys[i] = key
		root.data[i] = value
		root.nKey += 1
		return true
	}

	child := root.children[i]

	// 子节点的key是满的，要split
	if child.nKey == 2*minDegree-1 {
		root.split(i, minDegree)
		if root.keys[i] < key {
			i += 1
		}
	}

	root.children[i].insert(key, value, minDegree)
	return true
}

// 从 root 节点内删除 key， root 必须是叶子节点
func (root *Node) remove(key rune) {
	position := search(root.keys, key)

	if position > -1 && root.isLeaf {

		for i := position; i < int(root.nKey)-1; i += 1 {
			root.keys[i] = root.keys[i+1]
			root.data[i] = root.data[i+1]
		}
		root.nKey -= 1
	}
}

// root 节点key个数必须比 minDegree - 1大，才能执行这个方法,
// 或者
func (root *Node) delete(key rune, tree *Tree) bool {
	if root.has(key) {
		if root.isLeaf {
			// 叶子节点，直接删除就完事儿了
			root.remove(key)
			return true

		}
		// 内部节点，分两种情况
		position := search(root.keys, key)

		// case 1: root.children[position] 或 root.children[position + 1] 至少
		//     有 1 个节点的 key 大于 minDegree - 1 个；

		node := root.children[position+1]
		if isTooShort(node, tree.minDegree) {
			node = root.children[position]
		}

		if isTooShort(node, tree.minDegree) == false {
			// node == root.children[position + 1]
			// 将 node.keys[0] 替换到 root.keys[position],
			// 删除 node.keys[0] 即可
			if node.keys[0] > key {
				root.keys[position] = node.keys[0]
				return node.delete(node.keys[0], tree)
			}
			// node == root.children[position],
			// 做类似的操作即可
			root.keys[position] = node.keys[node.nKey-1]
			return node.delete(node.keys[node.nKey-1], tree)

		}

		// case2: root.children[position] 和 root.children[position + 1] 的 key 都是 minDegree - 1 个，
		//   将 root.children[position] root.children[position + 1] root.keys[position] 合并

		leftChild := node
		rightChild := root.children[position+1]
		i := leftChild.nKey

		// 合并 root.keys[position] 到 leftChild
		leftChild.keys[i] = root.keys[position]
		leftChild.nKey += 1
		i += 1

		// 将 rightChild 合并到 leftChild
		for j := uint(0); j < rightChild.nKey; j += 1 {
			leftChild.keys[i+j] = rightChild.keys[j]
			leftChild.data[i+j] = rightChild.data[j]
			if (leftChild.isLeaf) {
				leftChild.children[i+j] = rightChild.children[j]
			    leftChild.nChild += 1
			}
			
			leftChild.nKey += 1
		}
		if (leftChild.isLeaf) {
			leftChild.children[i+rightChild.nKey] = rightChild.children[rightChild.nKey]
			leftChild.nChild += 1
		}

		// 调整 root
		for i := position; i < int(root.nKey) - 1; i += 1 {
			root.keys[i] = root.keys[i+1]
			root.children[i+1] = root.children[i+1+1]
		}
		root.nKey -= 1
		root.nChild -= 1

		// 此时root可能是树的根节点，如果它没有key了，
		// 需要从新设置 root
		if tree.root == root && root.nKey == 0{
			tree.root = leftChild
		}

		// rightChild 不用管理了，等待GC回收

		return leftChild.delete(key, tree)
	}

	// 删除一个不存在的key，默认成功，因为这不会影响到什么
	if root.isLeaf {
		return true
	}
	nodeIndex := root.childIndexHasKey(key)
	if nodeIndex == -1 {
		return true
	}

	// 包含 key 的 root 的子树节点
	node := root.children[nodeIndex]

	// node的 key 数量比 minDegree 少，需要调整，以防删除key后，发生key个数下溢
	if isTooShort(node, tree.minDegree) {
		// node 的相邻节点，该节点的 key 必须多于 minDegree - 1 个
		var leftNeighbor *Node
		var rightNeighbor *Node

		// 只有右邻点
		if nodeIndex == 0 {
			rightNeighbor = root.children[1]
		} else if nodeIndex == int(root.nChild)-1 {
			// 只有左邻点
			leftNeighbor = root.children[root.nChild-2]
		} else {
			// 左右都有
			leftNeighbor, rightNeighbor = root.children[nodeIndex-1], root.children[nodeIndex+1]
		}

		neighbor := leftNeighbor

		if leftNeighbor == nil || isTooShort(leftNeighbor, tree.minDegree) {
			neighbor = rightNeighbor
			if rightNeighbor == nil || isTooShort(rightNeighbor, tree.minDegree) {
				neighbor = nil
			}
		}

		if neighbor != nil {
			// neighbor 是 node 的右邻点
			if neighbor.keys[0] > node.keys[node.nKey-1] {
				node.keys[node.nKey] = root.keys[nodeIndex]
				node.data[node.nKey] = root.data[nodeIndex]
				node.children[node.nKey+1] = root.children[nodeIndex+1]
				node.nKey += 1
				node.nChild += 1

				i := 0
				for ; i < int(root.nKey)-1; i += 1 {
					root.keys[i] = root.keys[i+1]
					root.data[i] = root.data[i+1]
					root.children[i+1] = root.children[i+2]
				}

				root.keys[i] = neighbor.keys[0]
				root.data[i] = neighbor.data[0]
				root.children[i+1] = root.children[i]
				root.children[i] = neighbor.children[0]

				i = 0
				for ; i < int(neighbor.nKey); i += 1 {
					neighbor.keys[i] = neighbor.keys[i+1]
					neighbor.data[i] = neighbor.data[i+1]
					neighbor.children[i+1] = neighbor.children[i+1+1]
				}
				neighbor.nChild -= 1
				neighbor.nKey -= 1
			} else {
				// neighbor 是 node 的左邻点
				i := node.nKey

				// node 节点内部右移
				for ; i > 0; i -= 1 {
					node.keys[i] = node.keys[i-1]
					node.data[i] = node.data[i-1]
					node.children[i+1] = node.children[i]
				}
				node.keys[i] = root.keys[root.nKey-1]
				node.data[i] = root.data[root.nKey-1]
				node.children[i] = root.children[root.nKey-1]
				node.nKey += 1
				node.nChild += 1

				// root 节点内部右移
				i = root.nKey - 1
				for ; i > 0; i -= 1 {
					root.keys[i] = root.keys[i-1]
					root.data[i] = root.data[i-1]
					root.children[i] = root.children[i-1]
				}
				root.children[root.nChild-1] = node

				// 将 neighbor 末尾key 和 child 转入 root
				i = neighbor.nKey - 1
				root.keys[0] = neighbor.keys[i]
				root.data[0] = neighbor.data[i]
				root.children[0] = root.children[1]
				root.children[1] = neighbor.children[neighbor.nChild-1]
				neighbor.nChild -= 1
				neighbor.nKey -= 1
			}

		} else {
			// neighbor == nil
			// 这种情况，node 左右无缘，没法借用一个key过来，
			// 只能和一个邻点合并

			if nodeIndex == 0 {
				neighbor = root.children[nodeIndex+1]
			} else if nodeIndex == int(root.nChild)-1 {
				// 这种情况下，确保下文取 key 的时候，不会发生溢出
				nodeIndex -= 1
				neighbor = root.children[nodeIndex]
			} else {
				neighbor = root.children[nodeIndex+1]
			}

			// 默认 neighbor 在 node 右侧， 如果不是这样，做交换处理，
			// 一律往左节点合并
			if neighbor.keys[0] <= node.keys[node.nKey-1] {
				neighbor, node = node, neighbor
			}

			// root节点的key 合并到 node 中
			i := node.nKey
			node.keys[i] = root.keys[nodeIndex]
			node.data[i] = root.data[nodeIndex]
			node.nKey += 1

			// neighbor 节点合并到 node 
			i += 1
			for j := uint(0); j < neighbor.nKey; j += 1 {
				node.keys[i+j] = neighbor.keys[j]
				node.data[i+j] = neighbor.data[j]
				node.children[i+j] = neighbor.children[j]
				node.nKey += 1
				node.nChild += 1
			}

			i = uint(nodeIndex)
			for ; i < root.nKey-1; i += 1 {
				root.keys[i] = root.keys[i+1]
				root.data[i] = root.data[i+1]
				root.children[i+1] = root.children[i+1+1]
			}
			root.nKey -= 1
			root.nChild -= 1

			if root == tree.root && root.nKey == 0 {
				tree.root = node
			}
		}

	}

	return node.delete(key, tree)

}

func (tree *Tree) Search(key rune) bool {
	result := tree.root.search(key)

	if result != nil {
		return true
	}

	return false
}

func (tree *Tree) Insert(key rune, value interface{}) bool {
	root := tree.root

	if isFull(root, tree.minDegree) {
		node := NewNode(tree.minDegree)

		node.children[0] = root

		node.split(0, tree.minDegree)

		tree.root = node

		node.isLeaf = false

		return node.insert(key, value, tree.minDegree)
	}

	return root.insert(key, value, tree.minDegree)
}

func (tree *Tree) Delete(key rune) bool {
	root := tree.root

	return root.delete(key, tree)
}

func (tree *Tree) Traverse() {
	root := tree.root

	var deal func(key rune)
	deal = func(key rune) {
		fmt.Printf("%c", key)
	}
    traverse(root, deal)
}


func main() {

	tree := NewTree(2)

	tree.Insert('a', 23)
	tree.Insert('d', 22)
	tree.Insert('t', 13)
	tree.Insert('h', 10)
	tree.Insert('i', 3)
    tree.Insert('p', 22)
    tree.Insert('y', 12)
	tree.Insert('w', 11)
	tree.Traverse()

	println()

	if tree.Search('h') {
		println("'a' exists")
	}

	tree.Delete('i')

	tree.Traverse()
}