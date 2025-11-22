package main

import "fmt"

func addNum(a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, fmt.Errorf("cannot enter 0 for both")
	}
	return a + b, nil
}

func subNum(a, b int) int {
	return a - b
}

func main() {

	a := 1
	b := 5

	c, err := addNum(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c, subNum(a, b))
}
