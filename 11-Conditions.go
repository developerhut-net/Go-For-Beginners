package main

import "fmt"


func main(){

/*
	if
	else
	else if
	switch
*/	

/*
	if condition {
		//...
	}

*/

if 40 > 20 {
	fmt.Println("40 is greater than 20")	
}

x := 20
y := 28

if x > y {
	fmt.Println("x is greater than y")	
} else if x < y {
	fmt.Println("y is greater than x")	
} else {
	fmt.Println("x & y are equal")	
}

num := 20

if num >= 10 {
	fmt.Println("num is more than 10")	
	if num > 15 {
		fmt.Println("num is more than 15")	
	}
} else {
	fmt.Println("num is less than 10")	
}

}