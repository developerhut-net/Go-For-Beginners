package main

import "fmt"


func main(){

	/*
		Print()
		Println()
		Printf()
	*/

	var i, j string = "Hello", "World";
	fmt.Print(i);
	fmt.Print(j);

	fmt.Print(i, "\n");
	fmt.Print(j, "\n");

	fmt.Print(i, "\n", j);
	fmt.Print(i, " ", j);

	fmt.Println(i, j);
	
	/*
		%v is used to print the value of the arguments
		%T is used to print the type of the arguments
	*/

	var x string = "Hello";
	var y int = 10;

	fmt.Printf("x has value: %v and type: %T \n", x, x);
	fmt.Printf("y has value: %v and type: %T", y, y);

}