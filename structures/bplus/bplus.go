package bplus

import (
	n "eda/structures/bplus/bplusNode"
	"fmt"
)

type Bplus[K any, V any] struct {
	Grade int
	Root  *n.BplusNode[K, V]

	CompareTo func(K, K) int

	FirstSheet *n.BplusNode[K, V]
}

func NewBplus[K any, V any](grade int, compareTo func(K, K) int) *Bplus[K, V] {
	root := &n.BplusNode[K, V]{
		Grade:     grade,
		Keys:      make([]*K, grade),
		Childs:    nil,
		Values:    make([]*V, grade),
		CompareTo: compareTo,
		Count:     0,
		Next:      nil,
	}

	return &Bplus[K, V]{
		Grade:      grade,
		Root:       root,
		CompareTo:  compareTo,
		FirstSheet: root,
	}
}

func (bplus *Bplus[K, V]) newInterNode() *n.BplusNode[K, V] {
	return &n.BplusNode[K, V]{
		Grade:     bplus.Grade,
		Keys:      make([]*K, bplus.Grade),
		Childs:    make([]*n.BplusNode[K, V], bplus.Grade+1),
		Values:    nil,
		CompareTo: bplus.CompareTo,
		Count:     0,
		Next:      nil,
	}
}

func (bplus *Bplus[K, V]) newSheetNode() *n.BplusNode[K, V] {
	return &n.BplusNode[K, V]{
		Grade:     bplus.Grade,
		Keys:      make([]*K, bplus.Grade),
		Childs:    nil,
		Values:    make([]*V, bplus.Grade),
		CompareTo: bplus.CompareTo,
		Count:     0,
		Next:      nil,
	}
}

func (bplus *Bplus[K, V]) Add(key K, value *V) {
	node := bplus.AddInNode(bplus.Root, key, value)

	if node != nil && node.Count == bplus.Grade {
		newK, newNode := bplus.Split(node)

		newRoot := bplus.newInterNode()
		newRoot.InsertInNode(newK, nil, newNode)
		newRoot.Childs[0] = bplus.Root

		bplus.Root = newRoot
	}
}

func (bplus *Bplus[K, V]) AddInNode(node *n.BplusNode[K, V], key K, value *V) *n.BplusNode[K, V] {
	if node.Childs == nil {
		node.InsertInNode(&key, value, nil)

		return node
	}

	pos := node.FindPos(key)

	cbn := bplus.AddInNode(node.Childs[pos], key, value)

	if cbn == nil {
		return nil
	}

	if cbn.Count == bplus.Grade {
		newK, newN := bplus.Split(cbn)
		node.InsertInNode(newK, nil, newN)

		return node
	}

	return nil
}

func (bplus *Bplus[K, V]) Split(node *n.BplusNode[K, V]) (*K, *n.BplusNode[K, V]) {
	if node.Childs == nil {
		return bplus.SplitSheet(node)
	}
	return bplus.SplitInter(node)
}

func (bplus *Bplus[K, V]) SplitSheet(node *n.BplusNode[K, V]) (*K, *n.BplusNode[K, V]) {
	newNode := bplus.newSheetNode()

	newNode.Next = node.Next
	node.Next = newNode

	middle := bplus.Grade / 2

	key := node.Keys[middle]

	for i := middle; i < bplus.Grade; i++ {
		newNode.Keys[i-middle] = node.Keys[i]
		node.Keys[i] = nil

		newNode.Values[i-middle] = node.Values[i]
		node.Values[i] = nil

		node.Count -= 1
		newNode.Count += 1
	}

	return key, newNode
}

func (bplus *Bplus[K, V]) SplitInter(node *n.BplusNode[K, V]) (*K, *n.BplusNode[K, V]) {
	newNode := bplus.newInterNode()

	middle := bplus.Grade / 2

	key := node.Keys[middle]
	node.Keys[middle] = nil

	newNode.Childs[0] = node.Childs[middle+1]
	node.Childs[middle+1] = nil
	node.Count -= 1

	for i := middle + 1; i < bplus.Grade; i++ {
		newNode.Keys[i-middle-1] = node.Keys[i]
		node.Keys[i] = nil

		newNode.Childs[i-middle] = node.Childs[i+1]
		node.Childs[i+1] = nil

		node.Count -= 1
		newNode.Count += 1
	}

	return key, newNode
}

func (bplus *Bplus[K, V]) Print() {
	fmt.Println("PRINT BPLUS")

	bplus.Root.PrintNode(0)
}

func (bplus *Bplus[K, V]) PrintLinkedList() {
	fmt.Println("PRINT LINKED LIST")

	if bplus.FirstSheet == nil {
		fmt.Println("NIL")
	}

	node := bplus.FirstSheet

	for {
		if node == nil {
			break
		}

		for i, v := range node.Values {
			if v == nil {
				break
			}
			fmt.Println(*node.Keys[i], *v)
		}

		node = node.Next
	}
}

func (bplus *Bplus[K, V]) GetArray() []V {
	res := make([]V, 0)

	if bplus.FirstSheet == nil {
		return res
	}

	node := bplus.FirstSheet

	for {
		if node == nil {
			break
		}

		for _, v := range node.Values {
			if v == nil {
				break
			}
			res = append(res, *v)
		}

		node = node.Next
	}

	return res
}
