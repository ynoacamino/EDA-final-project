package nodelist

type Node[T any] struct {
	data     *T
	nextNode *Node[T]
}

func (node *Node[T]) GetData() *T {
	return node.data
}

func (node *Node[T]) SetData(data *T) {
	node.data = data
}

func (node *Node[T]) GetNext() *Node[T] {
	return node.nextNode
}

func (node *Node[T]) SetNext(nextNode *Node[T]) {
	node.nextNode = nextNode
}

func NewNode[T any](data *T) *Node[T] {
	node := Node[T]{
		data:     data,
		nextNode: nil,
	}

	return &node
}
