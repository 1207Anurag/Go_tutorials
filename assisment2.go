package main

import (
	"fmt"
	
)

//Iota practie


func main() {
	const (
		january Month = 1+ iota
		february
		march
		april
		may
		june
		july
		august
		september
		october
		nov
		dec
	)
	fmt.Println(january,Month)
}
