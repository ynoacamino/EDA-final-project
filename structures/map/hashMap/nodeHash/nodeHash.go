package nodehash

type NodeHash[K any, V any] struct {
	key   K
	value *V
}

func (node *NodeHash[K, V]) GetKey() K {
	return node.key
}

func (node *NodeHash[K, V]) GetValue() *V {
	return node.value
}

func (node *NodeHash[K, V]) SetValue(value *V) {
	node.value = value
}

func NewNodeHash[K any, V any](key K, value *V) *NodeHash[K, V] {
	return &NodeHash[K, V]{
		key:   key,
		value: value,
	}
}
