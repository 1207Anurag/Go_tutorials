package main

import "fmt"

type Students struct {
	fname   string
	lname   string
	age     int
	gender  string
	disable bool
}
type Teacher struct {
	fname   string
	lname   string
	age     int
	subject string
	gender  string
}
type Course struct {
	Course_name     string
	Course_credit   int
	Course_duration int
}
type Details struct {
	candidate Students
	mentor    Teacher
	subject   Course
}

func main() {
	//assigning all the values into the result
	result := Details{
		candidate: Students{"Anurag", "Chaubey", 22, "male", false},
		mentor:    Teacher{"Ajay", "yadav", 24, "golang", "male"},
		subject:   Course{"golang", 20, 60},
	}
	fmt.Println(result)
}
