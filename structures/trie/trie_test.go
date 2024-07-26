package trie

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	trie := NewTrie[int]()

	value1 := 1

	trie.Add("Love Lost", &value1)
	trie.Add("C.R.E.A.M", &value1)

	trie.Print()

	suggests := trie.Suggest("holas")
	if suggests == nil {
		fmt.Println("No hay sugerencias")
		return
	}
	suggests.ForEach(func(value *int, i int) {
		fmt.Printf("Sugerencia: %d\n", *value)
	})

	for i := 0; i < 100; i++ {

	}
}

func TestAdd(t *testing.T) {
	trie := NewTrie[int]()

	value1 := 1
	value2 := 2
	value3 := 3

	trie.Add("hola", &value1)
	trie.Add("palabra", &value2)
	trie.Add("holaa1", &value3)

	if !trie.Search("hola") {
		t.Fatalf("No se encontró la palabra 'hola'")
	}

	if !trie.Search("palabra") {
		t.Fatalf("No se encontró la palabra 'palabra'")
	}

	if !trie.Search("holaa1") {
		t.Fatalf("No se encontró la palabra 'holaa1'")
	}
}

func TestGet(t *testing.T) {
	trie := NewTrie[int]()

	value1 := 1
	value2 := 2

	trie.Add("hola", &value1)

	if trie.Get("hola").Size() != 1 {
		t.Fatalf("Debería haber un solo valor")
	}

	if trie.Get("hola").Get(0) != &value1 {
		t.Fatalf("El valor no es el esperado")
	}

	trie.Add("hola", &value2)

	if trie.Get("hola").Size() != 2 {
		t.Fatalf("Debería haber dos valores")
	}

	if trie.Get("hola").Get(0) != &value2 {
		t.Fatalf("El valor no es el esperado")
	}
}

func TestSearch(t *testing.T) {
	trie := NewTrie[int]()

	value1 := 1
	value2 := 2
	value3 := 3

	trie.Add("valoruno", &value1)
	trie.Add("valordos", &value2)
	trie.Add("valortres", &value3)

	if !trie.Search("valoruno") {
		t.Fatalf("No se encontró la palabra 'valoruno'")
	}

	if !trie.Search("valordos") {
		t.Fatalf("No se encontró la palabra 'valordos'")
	}

	if !trie.Search("valortres") {
		t.Fatalf("No se encontró la palabra 'valortres'")
	}

}

func TestRemove(t *testing.T) {
	trie := NewTrie[int]()

	value1 := 1
	value2 := 2
	value3 := 3

	trie.Add("valoruno", &value1)
	trie.Add("valordos", &value2)
	trie.Add("valortres", &value3)

	if !trie.Search("valoruno") {
		t.Fatalf("No se encontró la palabra 'valoruno'")
	}

	trie.Remove("valoruno")

	if trie.Search("valoruno") {
		t.Fatalf("Se encontró la palabra 'valoruno'")
	}
}

func TestSuggest(t *testing.T) {
	trie := NewTrie[int]()

	value1 := 1
	value2 := 2
	value3 := 3

	trie.Add("valoruno", &value1)
	trie.Add("valordos", &value2)
	trie.Add("valortres", &value3)

	suggests := trie.Suggest("valo")

	if suggests == nil {
		t.Fatalf("No hay sugerencias")
	}
}
