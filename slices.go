package main

import (
	"fmt"
)

func main() {
	/*var s []int
	printSlices(s)
	//appending
	s = append(s, 4)
	printSlices(s)
	fmt.Println(s)
	s = append(s, 1, 2, 3, 8)
	fmt.Println(s)

	/*slices of slices
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	//players can take turn
	//slicing operation
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		//strings.Join() Function in Golang concatenates all the
		//elements present in the slice of string into a single string. This function is available in the string package.
		fmt.Println(strings.Join(board[i], " "))
	}*/
	//creation of slices using inbuilt function make
 //func make([]T, len,cap) []T
//  b:=[]string{"g","o","l","a","n","g"}
//  sliced:=b[1:4]
//  fmt.Println(sliced)

//appending at the end of the slice
// sliceArr:=[]int{1,2,3,4,5}
// fmt.Println(sliceArr,len(sliceArr),cap(sliceArr))
// sliceArr = append(sliceArr, 8,9,4)
// fmt.Println(sliceArr,len(sliceArr),cap(sliceArr))


}
func printSlices(s []int) {
	fmt.Println(len(s), cap(s))
}

