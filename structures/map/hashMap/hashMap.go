package hashmap

import (
	l "eda/structures/list/linkedList"
	city "eda/structures/map/hashFunc/cityHash"
	n "eda/structures/map/hashMap/nodeHash"
)

type HashMap[V any] struct {
	table []*l.LinkedList[n.NodeHash[string, V]] // slice de linkedList de nodeHash
	size  int

	capacity int

	loadFactor float32

	Colliding int
}

func NewHashMap[V any](capacity int) *HashMap[V] {
	const (
		LOAD_FACTOR float32 = 0.75
		SIZE        int     = 0
	)

	return &HashMap[V]{
		table:      make([]*l.LinkedList[n.NodeHash[string, V]], capacity),
		size:       SIZE,
		capacity:   capacity,
		loadFactor: LOAD_FACTOR,
	}
}

func (hashMap *HashMap[V]) rehash() {
	hashMap.capacity = hashMap.capacity * 2

	oldTable := hashMap.table

	hashMap.table = make([]*l.LinkedList[n.NodeHash[string, V]], hashMap.capacity)

	hashMap.size = 0
	hashMap.Colliding = 0

	for _, slot := range oldTable {
		if slot == nil {
			continue
		}

		slot.ForEach(func(node *n.NodeHash[string, V], i int) {
			hashMap.Put(node.GetKey(), node.GetValue())

		})
	}
}

func (hashmap *HashMap[V]) checkLoadFactor() {
	if (float32(hashmap.size) / float32(hashmap.capacity)) > hashmap.loadFactor {
		hashmap.rehash()
	}
}

func (hashMap *HashMap[V]) Put(key string, value *V) *V {
	nodeHash := n.NewNodeHash(key, value)

	hashedKey := city.CityHash32([]byte(key)) % uint32(hashMap.capacity)

	if hashMap.table[hashedKey] == nil {
		hashMap.table[hashedKey] = l.NewLinkedList(func(a, b n.NodeHash[string, V]) bool {
			return a.GetKey() == b.GetKey()
		})
	} else {
		hashMap.Colliding = hashMap.Colliding + 1
	}

	list := hashMap.table[hashedKey]

	currentNode := list.GetFirstNode()

	for currentNode != nil {
		if currentNode.GetData().GetKey() == key {
			oldValue := currentNode.GetData().GetValue()
			currentNode.GetData().SetValue(value)

			return oldValue
		}

		currentNode = currentNode.GetNext()
	}

	list.AddFirst(nodeHash)

	hashMap.size = hashMap.size + 1

	hashMap.checkLoadFactor()

	return value
}

func (hashmap *HashMap[V]) Get(key string) *V {
	hashedKey := city.CityHash32([]byte(key)) % uint32(hashmap.capacity)

	slot := hashmap.table[hashedKey]

	if slot == nil {
		return nil
	}

	currentNode := slot.GetFirstNode()

	for currentNode != nil {
		if currentNode.GetData().GetKey() == key {
			return currentNode.GetData().GetValue()
		}

		currentNode = currentNode.GetNext()
	}

	return nil
}

func (hashMap *HashMap[V]) Remove(key string) *V {
	hashedKey := city.CityHash32([]byte(key)) % uint32(hashMap.capacity)

	matchNode := n.NewNodeHash[string, V](key, nil)

	slot := hashMap.table[hashedKey]

	oldValue := slot.RemoveMatch(matchNode)

	if oldValue != nil {
		hashMap.size = hashMap.size - 1
	}

	return oldValue.GetValue()
}

func (hashMap *HashMap[V]) ContainsKey(key string) bool {
	hashedKey := city.CityHash32([]byte(key)) % uint32(hashMap.capacity)

	matchNode := n.NewNodeHash[string, V](key, nil)

	slot := hashMap.table[hashedKey]

	return slot.Contains(matchNode)
}

func (hashMap *HashMap[V]) IsEmpty() bool {
	return hashMap.size == 0
}

func (hashMap *HashMap[V]) Size() int {
	return hashMap.size
}

func (hasMap *HashMap[V]) Clear() {
	hasMap.table = make([]*l.LinkedList[n.NodeHash[string, V]], hasMap.capacity)
	hasMap.size = 0
}

func (hashMap *HashMap[V]) GetCapacity() int {
	return hashMap.capacity
}

func (hashMap *HashMap[V]) GetBalance() float32 {
	return float32(hashMap.size) / float32(hashMap.capacity)
}

func (hashMap *HashMap[V]) GetTable() []*l.LinkedList[n.NodeHash[string, V]] {
	return hashMap.table
}

// crea un forEach
func (hashMap *HashMap[V]) ForEach(f func(*n.NodeHash[string, V], int)) {
	for _, slot := range hashMap.table {
		if slot == nil {
			continue
		}

		slot.ForEach(f)
	}
}
