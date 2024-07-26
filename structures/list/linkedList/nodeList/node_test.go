package nodelist

import (
	"testing"
)

func TestNewNodeMethod(t *testing.T) {
	num := 2

	node := NewNode(&num)

	if node == nil {
		t.Fatal("Node struct not iniziatizated")
	}
}

func TestGetDataMethod(t *testing.T) {
	num := 2

	node := NewNode(&num)

	if *node.GetData() != 2 {
		t.Fatalf("The get data method not working, must be 2, but is %d", *node.GetData())
	}
}

func TestSetDataMethod(t *testing.T) {
	input := 2

	node := NewNode(&input)

	num := 3

	node.SetData(&num)

	if *node.GetData() != 3 {
		t.Fatalf("The set data method not working, must be 3, but is %d", *node.GetData())
	}
}

func TestGetNextMethod(t *testing.T) {
	num := 2

	node := NewNode(&num)

	if node.GetNext() != nil {
		t.Fatalf("The get next method not working, must be nil, but is %v", node.GetNext())
	}
}

func TestSetNextMethod(t *testing.T) {
	input1 := 2
	input2 := 3

	node := NewNode(&input1)
	node2 := NewNode(&input2)

	node.SetNext(node2)

	if node.GetNext() != node2 {
		t.Fatalf("The set next method not working, must be node2, but is %v", node.GetNext())
	}
}
