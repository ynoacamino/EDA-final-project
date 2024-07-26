package bplusnode

import "fmt"

type BplusNode[K any, V any] struct {
	Grade  int
	Keys   []*K
	Childs []*BplusNode[K, V]
	Values []*V

	Count int

	Next *BplusNode[K, V]

	CompareTo func(K, K) int
}

func (bplusNode *BplusNode[K, V]) FindPos(key K) int {
	i := 0

	for {
		if i == bplusNode.Count {
			break
		}
		current := bplusNode.Keys[i]

		if current == nil {
			break
		}

		if bplusNode.CompareTo(key, *current) <= 0 {
			break
		}

		i++
	}

	return i
}

func (bplusNode *BplusNode[K, V]) InsertInNode(key *K, value *V, rigthChild *BplusNode[K, V]) *V {
	if bplusNode.Childs == nil {
		return bplusNode.InsertInSheet(key, value)
	}

	bplusNode.InsertInInter(key, rigthChild)
	return value
}

func (bplusNode *BplusNode[K, V]) InsertInSheet(key *K, value *V) *V {
	pos := bplusNode.FindPos(*key)

	i := bplusNode.Count - 1

	for ; i >= pos; i-- {
		bplusNode.Keys[i+1] = bplusNode.Keys[i]
		bplusNode.Values[i+1] = bplusNode.Values[i]
	}

	bplusNode.Keys[pos] = key
	bplusNode.Values[pos] = value

	bplusNode.Count++
	return value
}

func (bplusNode *BplusNode[K, V]) InsertInInter(key *K, node *BplusNode[K, V]) *K {
	pos := bplusNode.FindPos(*key)

	i := bplusNode.Count - 1

	for ; i >= pos; i-- {
		bplusNode.Keys[i+1] = bplusNode.Keys[i]
		bplusNode.Childs[i+2] = bplusNode.Childs[i+1]
	}

	bplusNode.Keys[pos] = key
	bplusNode.Childs[pos+1] = node

	bplusNode.Count++
	return key
}

func (bplusNode *BplusNode[K, V]) PrintNode(tabs int) {
	tabsStr := ""
	for i := 0; i < tabs; i++ {
		tabsStr += "\t"
	}

	fmt.Println("------------------------", bplusNode.Count)
	for i := 0; i < bplusNode.Count; i++ {
		if bplusNode.Keys[i] == nil {
			break
		}
		fmt.Print(tabsStr, *bplusNode.Keys[i])
		if bplusNode.Values != nil {
			fmt.Print(" -> ", *bplusNode.Values[i])
		}
		fmt.Println()
	}
	if bplusNode.Childs != nil {
		for i := 0; i < bplusNode.Count; i++ {
			if bplusNode.Keys[i] == nil {
				break
			}
			bplusNode.Childs[i].PrintNode(tabs + 1)
		}
		bplusNode.Childs[bplusNode.Count].PrintNode(tabs + 1)
	}
}
