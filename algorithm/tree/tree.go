package tree

type Tree struct {
	root *Node
}

type Node struct {
	Left  *Node
	Right *Node
	Value interface{}
}

func NewTree(str interface{}) *Tree {
	return &Tree{
		root: &Node{Value: str},
	}
}

func (t *Tree) AppendLeft(node *Node, data interface{}) {
	if node == nil {
		return
	}
	node.Left = &Node{Value: data}
}

func (t *Tree) AppendRight(node *Node, data interface{}) {
	if node == nil {
		return
	}
	node.Right = &Node{Value: data}
}

// PrevTraverse 前序遍历
func (t *Tree) PrevTraverse(node *Node) {
	if node == nil {
		return
	}
	print("执行业务逻辑")
	t.PrevTraverse(node.Left)
	t.PrevTraverse(node.Right)
}

// MiddleTraverse 中序遍历
func (t *Tree) MiddleTraverse(node *Node) {
	if node == nil {
		return
	}
	t.MiddleTraverse(node.Left)
	print("执行业务逻辑")
	t.MiddleTraverse(node.Right)
}

// NextTraverse 后序遍历
func (t *Tree) NextTraverse(node *Node) {
	if node == nil {
		return
	}
	t.NextTraverse(node.Left)
	t.NextTraverse(node.Right)
	print("执行业务逻辑")
}
