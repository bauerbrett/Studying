package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	n := 10 // Example input
	fmt.Printf("Fibonacci number at position %d is %d\n", n, fibonacci(n))
}

//This implementation of the Fibonacci sequence has a time complexity of O(2^n),
//which is exponential in the size of the input. This is because each call to fibonacci(n) results in two additional recursive calls to fibonacci(n-1) and fibonacci(n-2).

//This algorithm for calculating the nth element of the Fibonacci sequence has a space complexity of , since it uses recursion and the call stack grows with the size of the input.
//Each recursive call requires additional memory to store the state of the function, and the total amount of memory required increases linearly with the size of the input.
