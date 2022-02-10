package main

import "fmt"

/*func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	prime := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(prime)

	//slicing
	var s []int = prime[2:5]
	fmt.Println(s)
}
*/
var pow = []int{1, 2, 8, 16, 32, 64, 128}

func main() {

}



/*func main() {
	//	arr1 := []int{4, 5, 6, 8, 12}

	//arr2 := []bool{true, false, true, false, true}
	r := []ArrayMerge{{4, true}, {5, false}, {6, true}, {8, false}, {12, true}}
	fmt.Println(r)
}*/

// type ArrayMerge struct {
// 	i int
// 	b bool
// }

func slicing(s []int) {

	slice1 := s[:0]
	fmt.Println(slice1)
	slice2 := s[:4]
	fmt.Println(slice2)
	slice3 := s[2:5]
	fmt.Println(slice3)
}
