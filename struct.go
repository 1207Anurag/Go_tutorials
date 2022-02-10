package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) person {
	p := person{name: name}
	p.age = 42
	return p

}
func main() {
	fmt.Println(person{"bob", 20})
	fmt.Println(person{"Anurag", 22})
	fmt.Println(person{"zopsmart", 10})
	fmt.Println(person{name: "kaanu"})
	fmt.Println(newPerson("rakshita"))
	c := Car{Name: "Ferrari", Model: "GTC4", Color: "Red", WeightInKg: 1920}
	fmt.Println("Car Name", c.Name)
	fmt.Println("car Model", c.Model)

}

type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}
