package main

import "fmt"

func main() {
	// f := func(i int) int { //anonymous function
	// 	fmt.Println("hello go", i)
	// 	return 0
	// }
	// func1(f) //f(10)  -> if g is not avilable

	fibo := fibonacci()
	fmt.Println(fibo(5))
}

// func func1(g func(int) int) {
// g(10)
// }

//explaination->
// we have declared an anonymous function which is accepting
//an integer and returning an integer and then we pass
//a function f to function1 , now g is also a function
//that accepts an function and returns and integer
//when we pass f as an argument to g, g in this current context will be equal to f
//means value of f is copeid to g, means g will be the
//function that we have declared in the main

func fibonacci() func(int) int {
	first, second := 0, 1
	return func(int) int {
		second, first = second+first, second
		return second
	}

}
