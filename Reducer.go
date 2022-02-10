package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}
	result := reduce(numbers, add)
	fmt.Println(result)
	result1 := reduce(numbers, sub)
	fmt.Println(result1)
	//r := AnotherWay(numbers)
	//fmt.Println(r)
}

func reduce(nums []int, method func(int, int) int) int {
	//these all are just function signature for add
	//and then we are calling in the method
	result := 0
	for _, element := range nums {
		result = method(result, element)
		//calling the method for every element, in the slice
	}
	return result
}
func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}

/*func AnotherWay(nums []int) int {
	result := 0
	for _, element := range nums {
		result += element
	}
	return result
}*/
