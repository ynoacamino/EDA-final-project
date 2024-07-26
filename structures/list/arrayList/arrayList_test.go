package arraylist

import (
	"fmt"
	"testing"
)

// Verifica que el constructor NewArrayList inicialice bien la lista.
func TestNewArrayList(t *testing.T) {
	arrayList := NewArrayList[int]()

	if arrayList == nil {
		t.Fatal("ArrayList struct no inicializada")
	} else {
		fmt.Println("TestNewArrayList: ArrayList inicializada correctamente")
	}
}

// Comprueba que el tamaño inicial de la lista sea 0.
func TestInitialSizeMethod(t *testing.T) {
	arrayList := NewArrayList[int]()

	if arrayList.Size() != 0 {
		t.Fatalf("El tamaño inicial de arrayList debe ser 0, pero es %d", arrayList.Size())
	} else {
		fmt.Println("TestInitialSizeMethod: Tamaño inicial es 0")
	}
}

// Verifica que el método Add agregue elementos correctamente y que el tamaño de la lista sea el esperado.
func TestAddMethod(t *testing.T) {
	arrayList := NewArrayList[int]()

	e1 := 0
	e2 := 1
	e3 := 2

	arrayList.Add(e1)
	arrayList.Add(e2)
	arrayList.Add(e3)

	if arrayList.Size() != 3 {
		t.Fatalf("El método Add no funciona correctamente, el tamaño debe ser 3 y es %d", arrayList.Size())
	} else {
		fmt.Println("TestAddMethod: Los elementos se agregaron correctamente")
	}
}

// Verifica que el método AddAt agregue elementos en los índices correctos.
func TestAddAtMethod(t *testing.T) {
	arrayList := NewArrayList[int]()

	e1 := 0
	e2 := 1
	e3 := 2

	arrayList.Add(e1)
	arrayList.Add(e2)

	// Intentar agregar e3 en el índice 1
	if err := arrayList.AddAt(1, e3); err != nil {
		t.Fatalf("El método AddAt devolvió un error: %s", err)
	}

	if value, err := arrayList.Get(1); err != nil || value != e3 {
		t.Fatalf("El método AddAt no funciona correctamente, el valor del índice 1 debe ser %d", e3)
	}
}

// Comprueba que el método Get obtenga los elementos correctos en los índices especificados.
func TestGetMethod(t *testing.T) {
	arrayList := NewArrayList[int]()

	e1 := 0
	e2 := 1

	arrayList.Add(e1)
	arrayList.Add(e2)

	if value, err := arrayList.Get(0); err != nil || value != e1 {
		t.Fatalf("El método Get no funciona correctamente, get(0) debe ser 0, pero es %d", value)
	} else {
		fmt.Println("TestGetMethod: El valor en el índice 0 es correcto")
	}

	if value, err := arrayList.Get(1); err != nil || value != e2 {
		t.Fatalf("El método Get no funciona correctamente en iteración, get(1) debe ser 1, pero es %d", value)
	} else {
		fmt.Println("TestGetMethod: El valor en el índice 1 es correcto")
	}
}

// Verifica que el método Contains funcione correctamente para encontrar elementos.
func TestContainsMethod(t *testing.T) {
	arrayList := NewArrayList[int]()

	e1 := 0
	e2 := 1

	arrayList.Add(e1)
	arrayList.Add(e2)

	if !arrayList.Contains(e1, func(a, b int) bool { return a == b }) {
		t.Fatalf("El método Contains no funciona, la lista debe contener 0")
	} else {
		fmt.Println("TestContainsMethod: La lista contiene el elemento 0")
	}

	if !arrayList.Contains(e2, func(a, b int) bool { return a == b }) {
		t.Fatalf("El método Contains no funciona, la lista debe contener 1")
	} else {
		fmt.Println("TestContainsMethod: La lista contiene el elemento 1")
	}
}

// Comprueba que el método IndexOf devuelva los índices correctos de los elementos.
func TestIndexOfMethod(t *testing.T) {
	arrayList := NewArrayList[int]()

	e1 := 0
	e2 := 1

	arrayList.Add(e1)
	arrayList.Add(e2)

	if index := arrayList.IndexOf(e1, func(a, b int) bool { return a == b }); index != 0 {
		t.Fatalf("El método IndexOf no funciona, el índice de 0 debe ser %d", index)
	} else {
		fmt.Println("TestIndexOfMethod: El índice del elemento 0 es correcto")
	}

	if index := arrayList.IndexOf(e2, func(a, b int) bool { return a == b }); index != 1 {
		t.Fatalf("El método IndexOf no funciona, el índice de 1 debe ser %d", index)
	} else {
		fmt.Println("TestIndexOfMethod: El índice del elemento 1 es correcto")
	}
}

// Verifica que el método Set reemplace correctamente los elementos en los índices especificados.
func TestSetMethod(t *testing.T) {
	arrayList := NewArrayList[int]()

	e1 := 0
	e2 := 1

	arrayList.Add(e1)
	arrayList.Add(e2)

	e3 := 2

	if err := arrayList.Set(0, e3); err != nil {
		t.Fatalf("El método Set devolvió un error: %s", err)
	} else {
		fmt.Println("TestSetMethod: El elemento en el índice 0 se reemplazó correctamente")
	}

	if value, err := arrayList.Get(0); err != nil || value != e3 {
		t.Fatalf("El método Set no funciona correctamente, el valor del índice 0 debe ser %d", value)
	} else {
		fmt.Println("TestSetMethod: El valor en el índice 0 es correcto después del set")
	}
}

// Comprueba que el método ForEach itere correctamente sobre los elementos de la lista y aplique una función en cada elemento.
func TestForEach(t *testing.T) {
	arrayList := NewArrayList[int]()

	e1 := 0
	e2 := 1

	arrayList.Add(e1)
	arrayList.Add(e2)

	arrayList.ForEach(func(data int, i int) {
		if i == 0 && data != e1 {
			t.Fatalf("El método ForEach no funciona, el valor del índice 0 debe ser %d", data)
		}

		if i == 1 && data != e2 {
			t.Fatalf("El método ForEach no funciona, el valor del índice 1 debe ser %d", data)
		}
	})

	fmt.Println("TestForEach: La iteración ForEach funciona correctamente")
}

// Prueba todos los métodos a la vez y muestra mensajes en cada test.
func TestAllMethods(t *testing.T) {
	fmt.Println("Iniciando TestAllMethods")

	arrayList := NewArrayList[int]()
	fmt.Println("ArrayList inicializada correctamente")

	e1 := 0
	e2 := 1
	e3 := 2

	arrayList.Add(e1)
	arrayList.Add(e2)
	arrayList.Add(e3)
	fmt.Println("Se agregaron tres elementos: 0, 1, 2")

	if arrayList.Size() != 3 {
		t.Fatalf("El tamaño de la lista debe ser 3 y es %d", arrayList.Size())
	} else {
		fmt.Println("El tamaño de la lista es correcto después de agregar elementos")
	}

	if value, err := arrayList.Get(0); err != nil || value != e1 {
		t.Fatalf("El valor en el índice 0 debe ser 0, pero es %d", value)
	} else {
		fmt.Println("El valor en el índice 0 es correcto")
	}

	if value, err := arrayList.Get(1); err != nil || value != e2 {
		t.Fatalf("El valor en el índice 1 debe ser 1, pero es %d", value)
	} else {
		fmt.Println("El valor en el índice 1 es correcto")
	}

	if !arrayList.Contains(e1, func(a, b int) bool { return a == b }) {
		t.Fatalf("La lista debe contener 0")
	} else {
		fmt.Println("La lista contiene el elemento 0")
	}

	if !arrayList.Contains(e2, func(a, b int) bool { return a == b }) {
		t.Fatalf("La lista debe contener 1")
	} else {
		fmt.Println("La lista contiene el elemento 1")
	}

	if index := arrayList.IndexOf(e1, func(a, b int) bool { return a == b }); index != 0 {
		t.Fatalf("El índice de 0 debe ser %d", index)
	} else {
		fmt.Println("El índice del elemento 0 es correcto")
	}

	if index := arrayList.IndexOf(e2, func(a, b int) bool { return a == b }); index != 1 {
		t.Fatalf("El índice de 1 debe ser %d", index)
	} else {
		fmt.Println("El índice del elemento 1 es correcto")
	}

	e4 := 3
	if err := arrayList.Set(1, e4); err != nil {
		t.Fatalf("El método Set devolvió un error: %s", err)
	} else {
		fmt.Println("El elemento en el índice 1 se reemplazó correctamente")
	}

	if value, err := arrayList.Get(1); err != nil || value != e4 {
		t.Fatalf("El valor en el índice 1 después del set debe ser 3, pero es %d", value)
	} else {
		fmt.Println("El valor en el índice 1 es correcto después del set")
	}

	arrayList.ForEach(func(data int, i int) {
		fmt.Printf("Elemento en el índice %d: %d\n", i, data)
	})

	fmt.Println("TestAllMethods: Todos los métodos se probaron correctamente")
}
