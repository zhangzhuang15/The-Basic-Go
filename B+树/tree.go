package main

// B+ 树主体特征和 B树 一样，区别在于：
// 1. 只有叶子结点存数据，其余节点只存索引key；
// 2. 同一个索引key，可能在叶子结点和非叶子结点各出现一次；
// 3. node.keys[i] ≥ node.children[i] 中所有的索尼 key，
// 同时，node.keys[i] 也存在于 node.children[i].keys 中
// 4. 叶子结点可以使用单向链表连接，也可以不用

type Data struct {
	values []interface{}
}

type Node struct {
	nKeys int
	keys []rune
	children []*Node
	right *Node
	isLeaf bool
    data *Data
}

type Tree struct {
	minDegree int
	root *Node
}

func NewNode(minDegree int, isLeaf bool) *Node{
	if isLeaf {
		return &Node{
			nKeys: 0,
			keys: make([]rune, 2 * minDegree - 1),
			children: nil,
			right: nil,
			isLeaf: true,
			data: &Data{ values: make([]interface{}, 2 * minDegree - 1)},
		}
	}

	return &Node{
		nKeys: 0,
		keys: make([]rune, 2 * minDegree - 1),
		children: make([]*Node, 2 * minDegree),
		right: nil,
		isLeaf: false,
		data: nil,
	}
}

// 返回值等于 key 的索引位置或者其后缀的索引位置
func bSearch(array []rune, _start, _end int , key rune) int {
	start, end := _start, _end

	for start < end {
		mid := int((start + end)/2)
		if array[mid] < key {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return end
}

// 从索引位置 position 开始，向右移动 keys， data， children 一位，
// 不更新 nkeys
func(root *Node) moveToRight(position int) {
	last:= root.nKeys

	// 先把末尾 child 右移
	if !root.isLeaf {
		root.children[last + 1] = root.children[last]
	}
	

	// 从 position 位置开始，到最后一个 key 右移一位，
	// 对应的 child 指针也右移
	for last -= 1; last >= position; last -= 1 {
		root.keys[last + 1] =  root.keys[last]
		
		// 叶子结点存储着数据，别忘记移动
		if root.isLeaf {
			root.data.values[last + 1] = root.data.values[last]
		} else {
			root.children[last + 1] = root.children[last]
		}
	}
}

func(root *Node) moveToLeft(position int) {
	start := position

	for ;start < root.nKeys; start += 1 {
		root.keys[start - 1] = root.keys[start]

		if root.isLeaf {
			root.data.values[start - 1] = root.data.values[start]
		} else {
			root.children[start - 1] = root.children[start]
		}
	}

	if !root.isLeaf {
		root.children[start - 1] = root.children[start]
	}
}

// 将 root.keys[start, end) 剪切到 node.keys[to, ?) 中，
// 相关的 data 和 children 也同步剪切, 更新 node 中的 nkeys，
// 不更新 root 中的 nkeys， 返回 root 中 减少的 key 数量，
// 这样设计的目的是，避免因为修改了 nkeys 值，影响到 moveToLeft 
// moveToRight
func(root *Node) clipTo(node *Node, start, end, to int)  int {
	i, j := start, to
	for ; i < end; i, j = i+1, j+1 {
		node.keys[j] = root.keys[i]
		if node.isLeaf {
			node.data.values[j] = root.data.values[i]
		} else {
			node.children[j] = root.children[i]
		}
		node.nKeys += 1
	}

	if !node.isLeaf {
		node.children[j] = root.children[i]
	}

	return i - start
}

// root.nkeys < 2 * minDegree - 1 ,
// root.children[i].nkeys == 2 * minDegree - 1,
// 将 root.children[i] 分裂
func(root *Node) split(i , minDegree int) {
	lastIndex := root.children[i].nKeys - 1
	splitPosition := int(lastIndex / 2)

	// 移动root的keys，为分裂腾出一个地方
	root.moveToRight(i)
	root.keys[i] = root.children[i].keys[splitPosition]
	root.nKeys += 1


	createLeafNode := root.children[i].isLeaf
	node := NewNode(minDegree, createLeafNode)
    root.children[i + 1] = node
    
	// 将分裂点右侧的key，children，data 转移到 node中,
	// splitPosition 位置留在原来的节点
	v := root.children[i].clipTo(node, splitPosition + 1, root.nKeys, 0)
	root.children[i].nKeys -= v

	// 单向链表调整
	if node.isLeaf {
		node.right = root.children[i].right
	    root.children[i].right = node
	}
	
}

// root.nKeys > minDegree - 1,
// root.children[i].nKeys == minDegree - 1,
// 将 root.children[i] 合并操作
func(root *Node) merge(i, minDegree int) {

	child := root.children[i]

	// 子节点处于最左边或者最右边，单独处理

	if i == 0 {
		right := root.children[i + 1]

		// 右结点 key 不够，合并
		if right.nKeys == minDegree - 1 {
			right.clipTo(child, 0, right.nKeys, child.nKeys)
		    root.children[i + 1] = child
		    root.moveToLeft(1)
		    root.nKeys -= 1
			// 单向链表处理
			if right.isLeaf {
				child.right = right.right
			}

		} else {
            root.keys[i] = right.keys[0]
			v := right.clipTo(child, 0, 1, child.nKeys)
			right.moveToLeft(1)
			right.nKeys -= v
		}
		
	} else if i == root.nKeys {
		left := root.children[i - 1]

		// 左结点key不够，合并
		if left.nKeys == minDegree - 1 {
			child.clipTo(left, 0, child.nKeys, left.nKeys)
		    root.nKeys -= 1

			// 单向链表处理
			if left.isLeaf {
				left.right = child.right
			}

		} else {
			child.moveToRight(0)
			v := left.clipTo(child, left.nKeys - 1, left.nKeys, 0)
			left.nKeys -= v
			root.keys[i - 1] = left.keys[left.nKeys - 1]
		}
		
	} else {
        left, right := root.children[i - 1], root.children[i + 1]

		if left.nKeys > minDegree - 1 {
			child.moveToRight(0)
			v := left.clipTo(child, left.nKeys - 1, left.nKeys, 0)
			left.nKeys -= v
			root.keys[i] = left.keys[left.nKeys - 1]

		} else if right.nKeys > minDegree - 1 {
			v := right.clipTo(child, 0, 1, child.nKeys)
            root.keys[i] = right.keys[0]
			right.moveToLeft(1)
			right.nKeys -= v

		} else {
			// 左右结点的key都不够，默认选择与左结点合并
			child.clipTo(left, 0, child.nKeys, left.nKeys)
		    root.moveToLeft(i)
		    root.nKeys -= 1
		    root.children[i-1] = left

			// 单向链表处理
			if left.isLeaf {
				left.right = child.right
			}
            
		}
		
	}
}

// root != nil
func(root *Node) has(key rune) bool {
	position := bSearch(root.keys, 0, root.nKeys, key)

	if position < root.nKeys && root.keys[position] == key {
		return true
	}

	if root.isLeaf {
		return false
	}

	return root.children[position].has(key)
}

// root != nil
func(root *Node) search(key rune) interface{} {
	position := bSearch(root.keys, 0, root.nKeys, key)

	if position < root.nKeys && root.keys[position] == key {
		if root.isLeaf {
			return root.data.values[position]
		}
	}

	if root.isLeaf {
		return nil
	}

	return root.children[position].search(key)
}

// root.nkeys < 2 * minDegree - 1
func(root *Node) insert(key rune, value interface{}, minDegree int) {
	position := bSearch(root.keys, 0, root.nKeys, key)

	if position < root.nKeys && root.keys[position] == key  {
		if root.isLeaf {
			root.data.values[position] = value
			return
		} 
        
		// root.children[position].insert(key, value, minDegree)
		// return
	}

	
	if root.isLeaf {
		root.moveToRight(position)
	    root.keys[position] = key
		root.data.values[position] = value
		root.nKeys += 1
		return
	} 

	if root.children[position].nKeys == 2 * minDegree - 1 {
		root.split(position, minDegree)
	}
	root.children[position].insert(key, value, minDegree)
}

// root.nkeys > minDegree - 1
func(root *Node) delete(key rune, minDegree int) *Node {
	position := bSearch(root.keys, 0, root.nKeys, key)

	if position < root.nKeys && root.keys[position] == key {
		if root.isLeaf {
			root.moveToLeft(position + 1)
			root.nKeys -= 1
			return root
		}

		// node := root.children[position].delete(key, minDegree)

		// if node != nil {
		// 	root.keys[position] = node.keys[node.nKeys - 1]
		// }

		// return root
	}

	if root.isLeaf {
		return nil
	}

	if root.children[position].nKeys == minDegree - 1 {
		root.merge(position, minDegree)
	}

	node := root.children[position].delete(key, minDegree)

	if node != nil {
		root.keys[position] = node.keys[node.nKeys - 1]
	}

	return root
}

func NewTree(minDegree int) *Tree {
	return &Tree {
		minDegree: minDegree,
		root: nil,
	}
}

func(tree *Tree) insert(key rune, value interface{}) {
	if tree == nil {
		return
	}
	if tree.root == nil {
		node := NewNode(tree.minDegree, true)
		node.keys[0] = key
		node.data.values[0] = value
		return
	}

	if tree.root.nKeys == 2 * tree.minDegree - 1 {
		node := NewNode(tree.minDegree, false)
		node.children[0] = tree.root
		node.split(0, tree.minDegree)
		node.insert(key, value, tree.minDegree)
		tree.root = node

	} else {
		tree.root.insert(key, value, tree.minDegree)
	}
}

func(tree *Tree) delete(key rune) {
	if tree == nil {
		return
	}
	if tree.root == nil {
		return
	}

	position := bSearch(tree.root.keys, 0, tree.root.nKeys, key)

	if tree.root.children[position].nKeys == tree.minDegree - 1 {
		tree.root.merge(position, tree.minDegree)
	}

	node := tree.root.delete(key, tree.minDegree)

	if node != nil {
		tree.root.keys[position] = node.keys[node.nKeys-1]
	}
}

func(tree *Tree) has(key rune) bool {
	if tree == nil {
		return false
	}
	if tree.root == nil {
		return false
	}

	return tree.root.has(key)
}

func(tree *Tree) get(key rune) interface{} {
	if tree == nil || tree.root == nil {
		return nil
	}

	return tree.root.search(key)
}

