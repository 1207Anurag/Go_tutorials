/*strconv
convert string to integer
string to boolean
string to float
*/
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	num := "10023"
	num2 := "true"
	num3 := "124.541"
	ConvertStringToInteger(num)
	ConvertStringToBoolean(num2)
	ConvertStringToFloat(num3)
	ConvertIntegerToString(543)
	ConvertFloatToString(3.14758)
	ConvertBooleanToString(false)
}
func ConvertStringToInteger(s string) {
	intVar, err := strconv.Atoi(s)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}

func ConvertStringToBoolean(s string) {
	b1, err := strconv.ParseBool(s)
	fmt.Println(b1, err, reflect.TypeOf(b1))
}

func ConvertStringToFloat(s string) {
	f, err := strconv.ParseFloat(s, 64)

	fmt.Println(f, err, reflect.TypeOf(f))
}
func ConvertIntegerToString(i int) {
	f := strconv.Itoa(i)
	fmt.Println(f, reflect.TypeOf(f))
}
func ConvertFloatToString(f float64) {
	var s string = strconv.FormatFloat(f, 'E', -1, 32)
	fmt.Println(s)
}
func ConvertBooleanToString(b bool) {
	var s string = strconv.FormatBool(b)
	fmt.Println(s)
}
