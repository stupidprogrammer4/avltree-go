package avltree

type NodeType interface {
	int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string
}

type Node[K NodeType] struct {
	key    K
	height int
	skew   int
	left   *Node[K]
	right  *Node[K]
	parent *Node[K]
}

func newNode[K NodeType](_key K) *Node[K] {
	return &Node[K]{
		key:    _key,
		height: 0,
		skew:   0,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

func (node *Node[K]) setKey(key K) {
	node.key = key
}

func (node *Node[K]) getKey() K {
	return node.key
}

func (node *Node[K]) setParent(parent *Node[K]) {
	node.parent = parent
}

func (node *Node[K]) getParent() *Node[K] {
	return node.parent
}

func (node *Node[K]) setLeft(left *Node[K]) {
	node.left = left
}

func (node *Node[K]) getLeft() *Node[K] {
	return node.left
}

func (node *Node[K]) setRight(right *Node[K]) {
	node.right = right
}

func (node *Node[K]) getRight() *Node[K] {
	return node.right
}

func (node *Node[K]) setHeight(height int) {
	node.height = height
}

func (node *Node[K]) getHeight() int {
	return node.height
}

func (node *Node[K]) setSkew(skew int) {
	node.skew = skew
}

func (node *Node[K]) getSkew() int {
	return node.skew
}
