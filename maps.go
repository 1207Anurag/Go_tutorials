package main

import "fmt"

func main() {
	//method 1
	var mp map[string]int = map[string]int{
		"apple":  5,
		"orange": 2,
		"banana": 8,
	}
	//method 2
	//maap := make(map[string]int) //this line will create a new empty map
	fmt.Println(mp["apple"])
	mp["apple"] = 900
	mp["Anurag"] = 777
	mp["sahil"] = 33

	delete(mp, "sahil")
	delete(mp, "Anurag")
	val, ok := mp["apple"]
	fmt.Println(val, ok)
	// if the key apple exist store the value in val, if
	//it does not exist check if yes okay will be set to true,
	//if not...ok will be false and val is specified as 0

	fmt.Println(mp)
	fmt.Println(map_practice(mp))
	//string is the key and int is the value
	//map dosent care about the order in which we have inserted the data

	//addding values inside the map happes very instantaneouly, quickly very faster than array
}
func map_practice(mp map[string]int) map[string]int {
	mp = map[string]int{
		"Anurag": 1,
		"Kruti":  2,
		"Monti":  3,
	}
	return mp
}
