package linkedlist

import (
	"errors"
	"fmt"

	nList "eda/structures/list/linkedList/nodeList"
)

type LinkedList[T any] struct {
	head    *nList.Node[T]
	length  int
	equeals func(a, b T) bool
}

func (list *LinkedList[T]) Size() int {
	return list.length
}

func (list *LinkedList[T]) Add(data *T) {
	node := nList.NewNode(data)

	if list.head == nil {
		list.head = node
		list.length = list.length + 1
		return
	}

	currentNode := list.head

	for currentNode.GetNext() != nil {
		currentNode = currentNode.GetNext()
	}

	currentNode.SetNext(node)

	list.length = list.length + 1
}

func (list *LinkedList[T]) AddFirst(data *T) {
	node := nList.NewNode(data)

	node.SetNext(list.head)

	list.head = node

	list.length = list.length + 1
}

func (list *LinkedList[T]) Get(n int) *T {
	if n < 0 || n > list.length {
		return nil
	}

	currentNode := list.head

	for i := 0; i < n; i++ {
		currentNode = currentNode.GetNext()
	}

	return currentNode.GetData()
}

func (list *LinkedList[T]) Remove(n int) *T {
	if n < 0 || n > list.length {
		return nil
	}

	currentNode := list.head

	if n == 0 {
		list.head = currentNode.GetNext()
		list.length = list.length - 1
		return currentNode.GetData()
	}

	count := 0
	for {
		if count+1 == n {
			saveNode := currentNode.GetNext()

			currentNode.SetNext(currentNode.GetNext().GetNext())
			list.length = list.length - 1

			return saveNode.GetData()
		}
		count = count + 1
	}
}

func (list *LinkedList[T]) RemoveMatch(data *T) *T {
	curentNode := list.head

	if list.equeals(*curentNode.GetData(), *data) {
		list.head = curentNode.GetNext()
		list.length = list.length - 1
		return curentNode.GetData()
	}

	for curentNode.GetNext() != nil {
		if list.equeals(*curentNode.GetNext().GetData(), *data) {
			saveNode := curentNode.GetNext()

			curentNode.SetNext(curentNode.GetNext().GetNext())
			list.length = list.length - 1

			return saveNode.GetData()
		}

		curentNode = curentNode.GetNext()
	}

	return nil
}

func (list *LinkedList[T]) GetFirstNode() *nList.Node[T] {
	return list.head
}

func (list *LinkedList[T]) Contains(data *T) bool {
	currentNode := list.head

	for currentNode != nil {
		if list.equeals(*currentNode.GetData(), *data) {
			return true
		}

		currentNode = currentNode.GetNext()
	}

	return false
}

func (list *LinkedList[T]) IndexOf(data *T) int {
	currentNode := list.head

	for i := 0; i < list.length && currentNode != nil; i++ {
		if list.equeals(*currentNode.GetData(), *data) {
			return i
		}
		currentNode = currentNode.GetNext()
	}
	return -1
}

func (list *LinkedList[T]) Set(n int, data *T) (*T, error) {
	if n < 0 || n > list.length {
		return nil, errors.New("index out of limits")
	}

	node := nList.NewNode[T](data)

	currentNode := list.head

	if n == 0 {
		node.SetNext(currentNode.GetNext())

		list.head = node
		return currentNode.GetData(), nil
	}

	count := 0
	for {
		if count+1 == n {
			saveNode := currentNode.GetNext()
			node.SetNext(currentNode.GetNext().GetNext())
			currentNode.SetNext(node)

			return saveNode.GetData(), nil
		}
		count = count + 1
	}
}

func NewLinkedList[T any](equals func(a, b T) bool) *LinkedList[T] {
	list := LinkedList[T]{
		head:    nil,
		length:  0,
		equeals: equals,
	}

	return &list
}

func (list *LinkedList[T]) ForEach(function func(*T, int)) {
	currentNode := list.head

	index := 0

	for currentNode != nil {
		function(currentNode.GetData(), index)

		currentNode = currentNode.GetNext()
		index = index + 1
	}
}

func (list *LinkedList[T]) Print() {
	currentNode := list.head

	if currentNode == nil {
		fmt.Println("Empty list")
		return
	}

	for currentNode != nil {
		fmt.Println(*currentNode.GetData())
		currentNode = currentNode.GetNext()
	}
}

func (list *LinkedList[T]) Parse() []T {
	if list == nil {
		return make([]T, 0)
	}

	currentNode := list.head

	if currentNode == nil {
		return nil
	}

	var result []T

	for currentNode != nil {
		result = append(result, *currentNode.GetData())
		currentNode = currentNode.GetNext()
	}

	return result
}
