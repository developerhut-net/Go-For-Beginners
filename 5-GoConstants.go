package main

import "fmt"


const (

	C int = 3;
	D string = "Hello";
	E = false;
)

func main(){

	// const

	// Syntax : const CONST_NAME type = value
	const PI float32 = 3.14;


	/*

		Type of constants
		1. Typed constants
		2. Untyped constants

	*/

	//typed 
	const A int = 1;

	//untyped
	const B = 5;

	fmt.Println(C);
	fmt.Println(D);
	fmt.Println(E);

}