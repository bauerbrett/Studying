package main

import "fmt"

type zigZag struct {
	v1, v2 []int
	i1, i2 int
	turn   bool
}

func initialize(v1, v2 []int) *zigZag {
	return &zigZag{v1, v2, 0, 0, true}
}

func (z *zigZag) next() (int, bool) {
	if z.turn {
		if len(z.v1) > 0 {
			front := z.v1[0]
			z.v1 = z.v1[1:]
			z.turn = false
			return front, true
		} else {
			z.turn = false

		}

	} else {
		if len(z.v2) > 0 {
			front := z.v2[0]
			z.v2 = z.v2[1:]
			z.turn = true
			return front, true
		} else {
			z.turn = true
		}

	}
	return 0, false
}
func (z *zigZag) hasNext() bool {
	return z.i1 < len(z.v1) || z.i2 < len(z.v2)
}

func main() {

	arr1 := []int{1, 2}
	arr2 := []int{3, 4, 5, 6}

	zigZag := initialize(arr1, arr2)

	for zigZag.hasNext() {
		val, exist := zigZag.next()
		if exist {
			fmt.Println(val)
		}

	}

}
