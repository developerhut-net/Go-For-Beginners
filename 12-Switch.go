package main

import "fmt"


func main(){


	/*
		Syntax

		switch expression {
		
		case a: 
			//code block
		case b: 
			//code block
		case c: 
			//code block
		...
		default:
			//code block
		}
	*/

	day := 1

	switch day {

	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("not valid day")

	}

	
	/*
		Syntax - Multi case 

		switch expression {
		
		case a,b: 
			//code block
		case c,d: 
			//code block
		case e,f: 
			//code block
		...
		default:
			//code block
		}
	*/

	switch day {

	case 1,3,5:
		fmt.Println("Odd weekday")
	case 2,4:
		fmt.Println("Even weekday")
	case 6,7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day number")
	}

}