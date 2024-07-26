package trieNode

import (
	l "eda/structures/list/linkedList"
	"fmt"
)

type TrieNode[T any] struct {
	key    byte
	values *l.LinkedList[T]
	childs []*TrieNode[T]
	end    bool
}

func NewTrieNode[T any](end bool, value *T, key byte) *TrieNode[T] {
	values := l.NewLinkedList(func(a, b T) bool {
		return &a == &b
	})

	if value != nil {
		values.AddFirst(value)
	}

	return &TrieNode[T]{
		childs: make([]*TrieNode[T], 27),
		end:    end,
		values: values,
		key:    key,
	}
}

func (node *TrieNode[T]) GetValues() *l.LinkedList[T] {
	return node.values
}

func (node *TrieNode[T]) GetInNode(bytes *[]byte, i int) *l.LinkedList[T] {
	key := (*bytes)[i]

	if node.childs[key] == nil {
		return nil
	}

	if len(*bytes) == i+1 {
		if node.childs[key].end {
			return node.childs[key].values
		} else {
			return nil
		}
	}

	return node.childs[key].GetInNode(bytes, i+1)
}

func (node *TrieNode[T]) GetKey() byte {
	return node.key
}

func (node *TrieNode[T]) AddInNode(bytes *[]byte, i int, value *T) {
	key := (*bytes)[i]
	isFinal := len(*bytes) == i+1

	if node.childs[key] == nil {
		if isFinal {
			node.childs[key] = NewTrieNode[T](true, value, key)
		} else {
			node.childs[key] = NewTrieNode[T](false, nil, key)
			node.childs[key].AddInNode(bytes, i+1, value)
		}
	} else {
		if isFinal {
			node.childs[key].end = true

			match := false

			node.childs[key].values.ForEach(func(v *T, i int) {
				if v == value {
					match = true
				}
			})

			if match {
				return
			}

			node.childs[key].values.AddFirst(value)
		} else {
			node.childs[key].AddInNode(bytes, i+1, value)
		}
	}
}

func (node *TrieNode[T]) SearchInNode(bytes *[]byte, i int) bool {
	// Chequear si el índice está dentro del rango del array de bytes
	if i >= len(*bytes) {
		return false
	}

	key := (*bytes)[i]

	if node.childs[key] == nil {
		return false
	}

	if len(*bytes) == i+1 && node.childs[key].end {
		return true
	}

	return node.childs[key].SearchInNode(bytes, i+1)
}

func (node *TrieNode[T]) RemoveInNode(bytes *[]byte, i int) *l.LinkedList[T] {
	// Chequear si el índice está dentro del rango del array de bytes
	if i >= len(*bytes) {
		return nil
	}

	key := (*bytes)[i]

	if node.childs[key] == nil {
		return nil
	}

	if len(*bytes) == i+1 && node.childs[key].end {
		node.childs[key].end = false

		values := node.childs[key].values
		node.childs[key].values = nil

		return values
	}

	return node.childs[key].RemoveInNode(bytes, i+1)
}

func (node *TrieNode[T]) SearchPreFix(bytes *[]byte, i int) *TrieNode[T] {
	key := (*bytes)[i]

	if node.childs[key] == nil {
		return nil
	}

	if len(*bytes) == i+1 {
		return node.childs[key]
	}

	return node.childs[key].SearchPreFix(bytes, i+1)
}

func (node *TrieNode[T]) GetAllChild(suggest *l.LinkedList[T]) {
	if node.end {
		if node.values == nil {
			fmt.Println("enconte el error")
			return
		}
		node.values.ForEach(func(v *T, i int) {
			suggest.AddFirst(v)
		})
	}

	for _, child := range node.childs {
		if child != nil {
			child.GetAllChild(suggest)
		}
	}
}

func (node *TrieNode[T]) Print(tab int) {
	strTab := ""
	for i := 0; i < tab; i++ {
		strTab += "\t"
	}
	fmt.Println(strTab, string(node.key+'a'), " | ", node.end)
	for _, child := range node.childs {
		if child != nil {
			child.Print(tab + 1)
		}
	}
}
