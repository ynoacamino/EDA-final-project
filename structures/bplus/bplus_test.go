package bplus

import "testing"

func TestNewPlus(t *testing.T) {
	bplus := NewBplus[int, int](30, func(i1, i2 int) int {
		if i1 == i2 {
			return 0
		} else if i1 > i2 {
			return 1
		} else {
			return -1
		}
	})

	e1 := 15
	bplus.Add(e1, &e1)

	e2 := 212
	bplus.Add(e2, &e2)

	e3 := 234
	bplus.Add(e3, &e3)

	e4 := 400
	bplus.Add(e4, &e4)

	e5 := 100
	bplus.Add(e5, &e5)

	e6 := 60
	bplus.Add(e6, &e6)

	e7 := 70
	bplus.Add(e7, &e7)

	e8 := 80
	bplus.Add(e8, &e8)

	e9 := 943
	bplus.Add(e9, &e9)

	e10 := 101
	bplus.Add(e10, &e10)

	bplus.PrintLinkedList()
}
