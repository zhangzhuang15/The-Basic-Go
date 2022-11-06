package main

type Data struct {
	values []interface{}
}

type Node struct {
	nKeys int
	keys []rune
	children []*Node
	isLeaf bool
    data *Data
	leaf *Node
}

type Tree struct {
	minDegree int
	root *Node
}

func bSearch(array []rune, key rune) int {
	start, end := 0, len(array)

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

func(root *Node) split()

func(root *Node) search(key rune) *Node {

	position := bSearch(root.keys, key) 

	if position < root.nKeys && root.keys[position] == key {
		return root
	}

	return root.children[position].search(key)
}

// root is not full
func(root *Node) insert(key rune, value interface{}, minDegree int) *Node {
	if root.isLeaf {
		position := bSearch(root.keys, key)

		if position < root.nKeys && root.keys[position] == key {
			root.data.values[position] = value
			return root
		}

		for end := root.nKeys; end > position; end -= 1 {
			root.keys[end] = root.keys[end - 1]
		}
		root.keys[position] = key

		for end := root.nKeys + 1; end > position; end -= 1 {
			root.data.values[end] = root.data.values[end - 1]
		}
		root.data.values[position] = value

		root.nKeys += 1

		return root
	}


	position := bSearch(root.keys, key)
        
	if position < root.nKeys && root.keys[position] == key {
		// if root.leaf is full, split it first
		node := root.leaf.insert(key, value, minDegree)
		root.leaf = node
		return root
	}

	r := root.children[position]

	// if r is full, split first
	leaf := r.insert(key, value, minDegree)

	r.leaf = leaf

	return leaf
}

func NewTree(minDegree int) *Tree {
	return &Tree {
		minDegree: minDegree,
		root: nil,
	}
}

