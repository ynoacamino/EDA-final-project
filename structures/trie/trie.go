package trie

import (
	l "eda/structures/list/linkedList"
	n "eda/structures/trie/trieNode"
	"strings"
)

type Trie[T any] struct {
	root *n.TrieNode[T]
}

func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{
		root: n.NewTrieNode[T](false, nil, 0),
	}
}

func (trie Trie[T]) Get(str string) *l.LinkedList[T] {
	bytes := convertToBytes(str)
	if len(bytes) == 0 {
		return nil
	}

	return trie.root.GetInNode(&bytes, 0)
}

func (trie Trie[T]) Add(str string, value *T) {
	bytes := convertToBytes(str)
	if len(bytes) == 0 {
		return
	}

	trie.root.AddInNode(&bytes, 0, value)
}

func (trie Trie[T]) Search(str string) bool {
	bytes := convertToBytes(str)
	if len(bytes) == 0 {
		return false
	}
	return trie.root.SearchInNode(&bytes, 0)
}

func convertToBytes(str string) []byte {
	str = strings.ToLower(str)
	var sb strings.Builder

	for i, l := range str {
		if l >= 'a' && l <= 'z' {
			sb.WriteRune(l - 'a')
		} else if l == 164 {
			sb.WriteRune('n' - 'a')
		}

		if i > (len(str)*3)/4 {
			break
		}
	}

	str = sb.String()
	if len(str) == 0 {
		return nil
	}
	return []byte(str)
}

func (trie *Trie[T]) Remove(str string) *l.LinkedList[T] {
	bytes := convertToBytes(str)
	return trie.root.RemoveInNode(&bytes, 0)
}

func (trie *Trie[T]) Suggest(str string) *l.LinkedList[T] {
	bytes := convertToBytes(str)
	if len(bytes) == 0 {
		return nil
	}
	node := trie.root.SearchPreFix(&bytes, 0)

	if node == nil {
		return nil
	}

	suggestSlice := l.NewLinkedList(func(a, b T) bool {
		return false
	})
	node.GetAllChild(suggestSlice)

	return suggestSlice
}

func (trie *Trie[T]) Print() {
	trie.root.Print(0)
}
