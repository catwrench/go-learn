package datastructures

type Bag[T any] struct {
	first *Node[T]
}

func NewBag[T any]() *Bag[T] {
	return &Bag[T]{}
}

func (b *Bag[T]) Add(data T) *Bag[T] {
	temp := NewNode[T](b.first.data)
	b.first.data = data
	b.first.next = temp
	return b
}

func (b *Bag[T]) Del() *Bag[T] {
	if b.IsEmpty() {
		return b
	}
	b.first = b.first.next
	return b
}

func (b *Bag[T]) IsEmpty() bool {
	return b.first == nil
}

func (b *Bag[T]) NewIterator() *Node[T] {
	return b.first
}
