package main

import "fmt"

func main() {
	p := []int{2, 3, 4, 5, 6, 7}
	evenNum := myfilter(p, isEven)
	fmt.Println(evenNum)
	OddNum := mymap(p, isOdd)
	fmt.Println(OddNum)
}

func mymap(a []int, fn func(int) int) []int {
	var filter []int
	for _, values := range a {
		if fn(values) == 1 {
			filter = append(filter, values)
		}
	}
	return filter
}

// misnomer - reduce
func myfilter(a []int, fn func(int) bool) []int {
	// TODO: implement
	var filter []int
	for _, values := range a {
		if fn(values) {
			filter = append(filter, values)
		}
	}
	return filter
}

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}
func isOdd(i int) int {
	if i%2 == 0 {
		return 0
	}
	return 1
}
