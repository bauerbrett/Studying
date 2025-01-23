package main

import "fmt"

type Solution struct{}

func (s *Solution) binaryConvert(n int) []string {

	nums := []int{}
	queue := []string{}

	for n > 0 {
		nums = append(nums, n)
		n -= 1
	}
	fmt.Println(nums)
	for i := len(nums) - 1; i >= 0; i-- {

		binarySlice := []rune{}
		stringSlice := []rune{}
		for nums[i] > 0 {
			remain := rune(nums[i]%2) + '0'
			nums[i] /= 2
			binarySlice = append(binarySlice, remain)

		}
		for len(binarySlice) > 0 { //While stack is not 0. THis is basically a while loop
			top := binarySlice[len(binarySlice)-1]         //Grab the top of stack
			binarySlice = binarySlice[:len(binarySlice)-1] // Pop the last element
			stringSlice = append(stringSlice, top)         //Append to the rune slice
		}

		queue = append(queue, string(stringSlice))
		//for i := 0; i < len(queue); i++ {
		//queue = append(queue, queue[i])
		//queue = queue[1:]
		//}

	}
	fmt.Println(queue)
	return nil

}

func main() {
	solution := &Solution{}

	dataInput := 2

	fmt.Println("Binary conversion:", solution.binaryConvert(dataInput))
}
