package main

import "fmt"

func main(){

/*
	1. Create/Declare a Function
	2. Call a Function
*/

/*
	Syntax
	func Function_Name() {
		//code block
	}

	func Function_Name(param1 type, param2 type, ...) {
		//code block
	}

	func Function_Name(param1 type, param2 type, ...) type {
		//code block

		return return_val
	}

*/

//my_function1(); //function call
//my_function1(); //function call
//my_function1(); //function call

/*
add(10, 20); // arguments
add(10, 30);
add(10, 40);

show(1, "John")
show("John", 8)
*/

total := add(1,2);
fmt.Println(total);

fmt.Println(myFunction(1, "David"));

a, b := myFunction(2, "John")
fmt.Println(a, b)

_, c := myFunction(9, "John")
fmt.Println(c)

d, _ := myFunction(25, "Jessy")
fmt.Println(d)

}

func show(num1 int, name string) {
	fmt.Printf("%v , %v", num1, name);
}

func myFunction(a int, name string) (addition int, greeting string) {

	addition = a + a;
	greeting = "Hello " + name;
	return;
}

func add(num1 int, num2 int) (num3 int) { // function Parameters
	num3 = num1 + num2
	return num3;
}

func add2(num1 int, num2 int) int { // function Parameters
	return num1 + num2;
}

func add1(num1 int, num2 int) { // function Parameters

	//var num1 int = 10;
	//var num2 int = 15;

	fmt.Printf("%v + %v = %v", num1, num2, num1+num2);

}

func My_Function1() {
	fmt.Println("Inside my function 1");
}

func my_function1() {
	fmt.Println("Inside my function 1");
}

/*
Naming Rules for Function
	1. Function name must start with letter
	2. Function name can contain only alpha-numeric chars & underscores
	3. Function names are case sensitive
	4. Function name can not contains spaces
*/