// This is single line comment
package main2

import "fmt"

var b string;

func main(){

/*
	int : 123 or -123
	float32 : 19.33 or -19.33
	string : "Hello World"
	bool : true / false
*/

//1. var keyword
//Syntax : var variable_name type
var name string;
var age int;


//2. := sign
//Syntax : variable_name := value
x := 2;


//Declear & initialize with value
var a int = 20;
y := 40;


//Value assignment
name = "John";

var a int;		//0
var b string;	// ""
var c bool;		// false

fmt.Println(a);
fmt.Println(b);
fmt.Println(c);

/*

var
1. can use this inside & outside the function
2. delcaration & value assignment can done seperately

:=
1. inside the function
2. delcaration & value assignment cannot done seperately
*/


}