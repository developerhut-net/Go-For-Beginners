package main

import "fmt"

func main(){

	//test(1);
	factorial(4);

}

// Formula: n! = n * (n-1) * (n-2) * ... * 2 * 1 
func factorial(x int) (y int) {
	
	if x > 0 {
		n := factorial(x-1);
		y = x * n;
		fmt.Printf("%v * %v = %v\n",x,n,y)
	} else {
		y = 1
	}

	return
}


func test(x int) int {

	if(x == 5) {
		return 0;
	}

	fmt.Println(x);
	return test(x + 1);

}