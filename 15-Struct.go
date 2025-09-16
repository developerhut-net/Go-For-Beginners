package main

import "fmt"

type Person struct {

	name string
	age int
	salary int
	designation string

}

func main(){

	//struct - custom / user defined data type
	//int, float32, bool, string

	/*
		type stuct_name struct {

			member1 type
			member2 type
			...
		}

	*/

	var p1 Person
	//var p2 Person

	p1.name = "John"
	p1.age = 28
	p1.salary = 20000
	p1.designation = "Developer"

	print_person(p1)

}

func print_person(p1 Person) {
	fmt.Printf("name : %v\n", p1.name);
	fmt.Printf("age : %v\n", p1.age);
	fmt.Printf("salary : %v\n", p1.salary);
	fmt.Printf("designation : %v\n", p1.designation);
}