package main

import "fmt"

var i, j int = 1, 2 //global variables

func swap(x, y string) (string, string) {
	return y, x
}

/*func main() {
	a, b := swap("Go", "Java")
	fmt.Println(a, b)
	var i int
	fmt.Println(split(17))
	fmt.Println(i, c, python, java)

}*/

//using naked return values
func split(sum int) (x, y int) { //x and y are the things that we have to return
	x = sum * 4 / 9
	y = sum - x
	return
}

/*func main() {
var c, python, java bool = true, false, true
fmt.Println(i, j, c, python, java)*/

/*func main() {
var i, j int = 1, 2
k := 3
c, python, java := true, false, "no!!"
fmt.Println(i, j, k, c, python, java)*/

var (
	ToBe    bool   = false
	MaxUnit uint64 = 1<<64 - 1
)

/*func main() {
	i := 42
	g := 0.867 + 0.5*float64(i) //type casting technique
	fmt.Println(g)
}
*/
//stacking defers
// func main() {
// fmt.Println("counting ")
// for i := 0; i < 10; i++ {
// defer fmt.Println(i)
// }
// fmt.Println("done")
// }

/*func main() {
	i, j := 42, 2701
	//if the value is directly printed it wll print the address of the p whereas if we use * before it will print the
	//value the the digit
	p := &i
	k := 45
	n := &k
	fmt.Println(n)
	fmt.Println(*n)
	fmt.Println(*p)
	fmt.Println(j)
}
*/
func main() {
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1} //Y: 0 is implicit
	v3 := Vertex{}     // x and y both are implicit
	p := &Vertex{1, 2}
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(*p)

}

type Vertex struct {
	X int
	Y int
}
