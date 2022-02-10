package main

import (
	"fmt"
)

func main() {
	var s []int
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

}
func printSlices(s []int) {
	fmt.Println(len(s), cap(s))
}
