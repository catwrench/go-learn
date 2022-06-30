package datastructures

// Queue 队列
type Queue[T any] struct {
	first *Node[T]
	last  *Node[T]
	num   int
	index int // 当前索引，用于支持迭代
}

type Node[T any] struct {
	data T
	next *Node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{
		data: data,
	}
}

// 向队列尾部添加元素
func (q *Queue[T]) EnQueue(data T) *Queue[T] {
	temp := q.last
	q.last = NewNode[T](data)
	q.last.next = nil
	if q.IsEmpty() {
		q.first = q.last
	} else {
		temp.next = q.last
	}
	q.num++
	return q
}

func (q *Queue[T]) DeQueue() *Queue[T] {
	// 队列先进先出，所以删除表头元素
	if q.IsEmpty() {
		return q
	}
	q.first = q.first.next
	q.num--
	return q
}

func (q *Queue[T]) IsEmpty() bool {
	return q.num == 0
}

func (q *Queue[T]) NewIterator() *Node[T] {
	q.index = q.num
	return q.first
}

func (n *Node[T]) HasNext() bool {
	// 因为是栈，所以是从最后一位向前遍历的
	return n.next != nil
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}
