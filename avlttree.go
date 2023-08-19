package avltree

type AVLTree[K NodeType] struct {
	root *Node[K]
}

func NewTree[K NodeType]() *AVLTree[K] {
	return &AVLTree[K]{
		root: nil,
	}
}

func max(h1, h2 int) int {
	if h1 > h2 {
		return h1
	}
	return h2
}

func (tree *AVLTree[K]) Insert(key K) {
	node := newNode[K](key)
	if tree.root == nil {
		tree.root = node
		return
	}

	tree.insertRecursive(tree.root, node)
}

func (tree *AVLTree[K]) Delete(key K) {
	if tree.root == nil {
		return
	}
	if tree.root.getLeft() == nil && tree.root.getRight() == nil && tree.root.getKey() == key {
		tree.root = nil
	}
	node := tree.find(tree.root, key)

	if node == nil {
		return
	}
	tree.deleteRecursive(node)
}

func (tree *AVLTree[K]) Find(key K) bool {
	return tree.find(tree.root, key) != nil
}

func (tree *AVLTree[K]) getMinNode(node *Node[K]) *Node[K] {
	if node.getLeft() == nil {
		return node
	}
	return tree.getMinNode(node.getLeft())
}

func (tree *AVLTree[K]) getMaxNode(node *Node[K]) *Node[K] {
	if node.getRight() == nil {
		return node
	}
	return tree.getMaxNode(node.getRight())
}

func (tree *AVLTree[K]) Max() K {
	return tree.getMaxNode(tree.root).getKey()
}

func (tree *AVLTree[K]) Min() K {
	return tree.getMinNode(tree.root).getKey()
}

func (tree *AVLTree[K]) insertRecursive(curr, node *Node[K]) {
	if node.getKey() < curr.getKey() {
		if curr.getLeft() == nil {
			node.setParent(curr)
			curr.setLeft(node)
			tree.fix(curr)
			return
		}
		tree.insertRecursive(curr.getLeft(), node)
	} else {
		if curr.getRight() == nil {
			node.setParent(curr)
			curr.setRight(node)
			tree.fix(curr)
			return
		}
		tree.insertRecursive(curr.getRight(), node)
	}
}

func (tree *AVLTree[K]) deleteRecursive(node *Node[K]) {
	if node.getLeft() == nil {
		right := node.getRight()
		tree.mixWithParent(node, right)
	} else if node.getRight() == nil {
		left := node.getLeft()
		tree.mixWithParent(node, left)
	} else {
		minNode := tree.getMinNode(node)
		node.setKey(minNode.getKey())
		tree.deleteRecursive(minNode)
	}
}

func (tree *AVLTree[K]) fix(node *Node[K]) {
	if node == nil {
		return
	}
	tree.update(node)
	if node.getSkew() == 2 {
		if node.getRight().getSkew() == -1 {
			tree.rotateRight(node.getRight())
		}
		tree.rotateLeft(node)
	} else if node.getSkew() == -2 {
		if node.getLeft().getSkew() == 1 {
			tree.rotateLeft(node.getLeft())
		}
		tree.rotateRight(node)
	}
	tree.fix(node.getParent())
}

func (tree *AVLTree[K]) rotateLeft(node *Node[K]) {
	right := node.getRight()
	A := node.getLeft()
	B := right.getLeft()
	C := right.getRight()

	x := node.getKey()
	y := right.getKey()

	node.setKey(y)
	right.setKey(x)

	node.setLeft(right)
	right.setParent(node)

	right.setLeft(A)
	if A != nil {
		A.setParent(right)
	}

	right.setRight(B)
	if B != nil {
		B.setParent(right)
	}

	node.setRight(C)
	if C != nil {
		C.setParent(node)
	}

	tree.update(right)
	tree.update(node)
}

func (tree *AVLTree[K]) rotateRight(node *Node[K]) {
	left := node.getLeft()
	A := left.getLeft()
	B := left.getRight()
	C := node.getRight()

	x := left.getKey()
	y := node.getKey()

	left.setKey(y)
	node.setKey(x)

	node.setRight(left)
	left.setParent(node)

	node.setLeft(A)
	if A != nil {
		A.setParent(node)
	}

	left.setLeft(B)
	if B != nil {
		B.setParent(left)
	}

	left.setRight(C)
	if C != nil {
		C.setParent(left)
	}

	tree.update(left)
	tree.update(node)
}

func (tree *AVLTree[K]) mixWithParent(node, child *Node[K]) {
	if node == tree.root {
		tree.root = child
		tree.root.setParent(nil)
		return
	}
	par := node.getParent()
	if node == par.getLeft() {
		par.setLeft(child)
	} else {
		par.setRight(child)
	}

	if child != nil {
		child.setParent(par)
	}
	tree.fix(par)
}

func (tree *AVLTree[K]) update(node *Node[K]) {
	right, left := tree.getHeight(node.getRight()), tree.getHeight(node.getLeft())
	node.setSkew(right - left)
	node.setHeight(max(left, right) + 1)
}

func (tree *AVLTree[K]) find(node *Node[K], key K) *Node[K] {
	if node == nil {
		return nil
	}
	if key < node.getKey() {
		return tree.find(node.getLeft(), key)
	} else if key > node.getKey() {
		return tree.find(node.getRight(), key)
	} else {
		return node
	}
}

func (tree *AVLTree[K]) getHeight(node *Node[K]) int {
	if node == nil {
		return -1
	}
	return node.getHeight()
}
