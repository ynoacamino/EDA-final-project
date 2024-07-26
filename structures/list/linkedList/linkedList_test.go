package linkedlist

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	linkedList := NewLinkedList(func(a, b int) bool {
		return a == b
	})

	if linkedList == nil {
		t.Fatal("LinkedList struct not iniziatizated")
	}
}

func TestInitialSizeMethod(t *testing.T) {
	linkedList := NewLinkedList(func(a, b int) bool {
		return a == b
	})

	if linkedList.Size() != 0 {
		t.Fatalf("The initial length of linkedList must be 0, but is %d", linkedList.Size())
	}
}

func TestAddMethod(t *testing.T) {
	linkedList := NewLinkedList(func(a, b int) bool {
		return a == b
	})

	e1 := 0
	e2 := 1
	e3 := 2

	linkedList.Add(&e1)
	linkedList.Add(&e2)
	linkedList.Add(&e3)

	if linkedList.Size() != 3 {
		t.Fatalf("The add method not working, the length must be 3 and is %d", linkedList.Size())
	}
}

func TestGetMethod(t *testing.T) {
	linkedList := NewLinkedList(func(a, b int) bool {
		return a == b
	})

	e1 := 0
	e2 := 1

	linkedList.Add(&e1)
	linkedList.Add(&e2)

	if *linkedList.Get(0) != e1 {
		t.Fatalf("The get method not working, get(0) must be 0, but is %d", *linkedList.Get(0))
	}

	if *linkedList.Get(0) != e1 || *linkedList.Get(1) != e2 {
		t.Fatalf("The get method not working in iteration, get(1) must be 1, but is %d", *linkedList.Get(1))
	}
}

func TestConteinsMethod(t *testing.T) {
	linkedList := NewLinkedList(func(a, b int) bool {
		return a == b
	})

	e1 := 0
	e2 := 1

	linkedList.Add(&e1)
	linkedList.Add(&e2)

	if !linkedList.Contains(&e1) {
		t.Fatalf("The contains method not working, the list must have 0")
	}

	if !linkedList.Contains(&e2) {
		t.Fatalf("The contains method not working, the list must have 1")
	}
}

func TestIndexOfMethod(t *testing.T) {
	linkedList := NewLinkedList(func(a, b int) bool {
		return a == b
	})

	e1 := 0
	e2 := 1

	linkedList.Add(&e1)
	linkedList.Add(&e2)

	if linkedList.IndexOf(&e1) != 0 {
		t.Fatalf("The indexOf method not working, the index of 0 must be %d", linkedList.IndexOf(&e1))
	}

	if linkedList.IndexOf(&e2) != 1 {
		t.Fatalf("The indexOf method not working, the index of 1 must be %d", linkedList.IndexOf(&e2))
	}
}

func TestSetMethod(t *testing.T) {
	linkedList := NewLinkedList(func(a, b int) bool {
		return a == b
	})

	e1 := 0
	e2 := 1

	linkedList.Add(&e1)
	linkedList.Add(&e2)

	e3 := 2

	linkedList.Set(0, &e3)

	if *linkedList.Get(0) != e3 {
		t.Fatalf("The set method not working, the value of index 0 must be %d", *linkedList.Get(0))
	}
}

func TestForEach(t *testing.T) {
	linkedList := NewLinkedList(func(a, b int) bool {
		return a == b
	})

	e1 := 0
	e2 := 1

	linkedList.Add(&e1)
	linkedList.Add(&e2)

	linkedList.ForEach(func(data *int, i int) {
		if i == 0 && *data != e1 {
			t.Fatalf("The forEach method not working, the value of index 0 must be %d", *data)
		}

		if i == 1 && *data != e2 {
			t.Fatalf("The forEach method not working, the value of index 1 must be %d", *data)
		}
	})
}
