package main

import "fmt"

func main() {

	// #1
	employee := make(map[string]string)

	fmt.Println(employee)
	//Add key,value
	employee["name"] = "bob"
	employee["title"] = "security engineer"
	//Fetch them
	fmt.Println(employee["name"], "and", employee["title"])
	//Delete
	delete(employee, "title")
	fmt.Println(employee)

	// #2
	array := []int{1, 3, 5, 7, 9}

	//Access by index
	fmt.Println(array[0])

	//Search Unsorted Array
	for x, y := range array {
		if y == 9 {
			fmt.Println("Found at index:", x)
		}
	}

	//Search Sorted Array
	sortArray := []int{10, 20, 30, 40, 50}
	target := 30
	low := 0
	high := len(sortArray) - 1

	for low <= high {
		mid := (low + high) / 2
		if sortArray[mid] == target {
			fmt.Println("Found at index:", mid)
			return // Exit after finding the target
		} else if sortArray[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Target not found")
}
