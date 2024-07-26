package hashmap

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	hashMap := NewHashMap[int](10)

	e1 := 1

	hashMap.Put("hola", &e1)

	mySlice := make([]int, 9)
	// slice de enteros de capacitadad 9, longitud 9 y elementos 0

	fmt.Println("Capacidad", cap(mySlice))
	fmt.Println("Longitud", len(mySlice))

	// foreach
	for i, v := range mySlice {
		fmt.Println(i, v)
	}

	//mySlice[9] = 10
	mySlice = append(mySlice, 10)

	fmt.Println("Capacidad", cap(mySlice))
	fmt.Println("Longitud", len(mySlice))

	fmt.Println(mySlice)
}

// crea test para cada funcion

// TestPut
func TestPut(t *testing.T) {
	hashMap := NewHashMap[int](10)

	e1 := 1

	hashMap.Put("hola", &e1)

	if *hashMap.Get("hola") != e1 {
		t.Fatalf("The put method not working, the value of key hola must be %d", 1)
	}
}

// TestRemove
func TestRemove(t *testing.T) {
	hashMap := NewHashMap[int](10)

	e1 := 1

	hashMap.Put("hola", &e1)

	if *hashMap.Remove("hola") != e1 {
		t.Fatalf("The remove method not working, the value of key hola must be %d", 1)
	}
}

// TestGet
func TestGet(t *testing.T) {
	hashMap := NewHashMap[int](10)

	e1 := 1

	hashMap.Put("hola", &e1)

	if *hashMap.Get("hola") != e1 {
		t.Fatalf("The get method not working, the value of key hola must be %d", 1)
	}
}

// TestSize
func TestSize(t *testing.T) {
	hashMap := NewHashMap[int](10)

	e1 := 1

	hashMap.Put("hola", &e1)

	if hashMap.Size() != 1 {
		t.Fatalf("The size method not working, the size must be %d", 1)
	}
}

// TestIsEmpty
func TestIsEmpty(t *testing.T) {
	hashMap := NewHashMap[int](10)

	if !hashMap.IsEmpty() {
		t.Fatalf("The isEmpty method not working, the hashMap must be empty")
	}
}

// TestContainsKey
func TestContainsKey(t *testing.T) {
	hashMap := NewHashMap[int](10)

	e1 := 1

	hashMap.Put("hola", &e1)

	if !hashMap.ContainsKey("hola") {
		t.Fatalf("The containsKey method not working, the key hola must be in the hashMap")
	}
}

// TestPutWithCollision
func TestPutWithCollision(t *testing.T) {
	hashMap := NewHashMap[int](10)

	e1 := 1
	e2 := 2

	hashMap.Put("hola", &e1)
	hashMap.Put("hola", &e2)

	if *hashMap.Get("hola") != e2 {
		t.Fatalf("The put method not working with collision, the value of key hola must be %d", 2)
	}
}

// TestRemoveWithCollision
func TestRemoveWithCollision(t *testing.T) {
	hashMap := NewHashMap[int](10)

	e1 := 1
	e2 := 2

	hashMap.Put("hola", &e1)
	hashMap.Put("hola", &e2)

	if *hashMap.Remove("hola") != e2 {
		t.Fatalf("The remove method not working with collision, the value of key hola must be %d", 2)
	}
}

// TestGetWithCollision
func TestGetWithCollision(t *testing.T) {
	hashMap := NewHashMap[int](10)

	e1 := 1
	e2 := 2

	hashMap.Put("hola", &e1)
	hashMap.Put("hola", &e2)

	if *hashMap.Get("hola") != e2 {
		t.Fatalf("The get method not working with collision, the value of key hola must be %d", 2)
	}
}

// TestSizeWithCollision
func TestSizeWithCollision(t *testing.T) {
	hashMap := NewHashMap[int](10)

	e1 := 1
	e2 := 2

	hashMap.Put("hola", &e1)
	hashMap.Put("hola", &e2)

	if hashMap.Size() != 1 {
		t.Fatalf("The size method not working with collision, the size must be %d but is %d", 1, hashMap.Size())
	}
}

// TestRehashing
func TestRehashing(t *testing.T) {
	hashMap := NewHashMap[int](3)

	e1 := 1
	e2 := 2
	e3 := 3

	hashMap.Put("hola", &e1)
	fmt.Println("LoadFactor 1 ", hashMap.GetBalance())

	hashMap.Put("hola1", &e2)
	fmt.Println("LoadFactor 2", hashMap.GetBalance())

	hashMap.Put("hola2", &e3)
	fmt.Println("LoadFactor 3", hashMap.GetBalance())

	// rehashing

	if hashMap.GetCapacity() != 6 {
		t.Fatalf("The rehashing method not working, the capacity must be %d but is %d", 6, hashMap.GetCapacity())
	}

	if hashMap.Size() != 3 {
		t.Fatalf("The rehashing method not working, the size must be %d", 3)
	}
}
