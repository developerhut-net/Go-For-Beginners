package main

import "fmt"


func main(){

	/*
		1. With the var keyword
		2. With the := sign
	*/

	/*
		1. With the var keyword
		
			var array_name = [length]data_type{values}
			or
			var array_name = [...]data_type{values}
	*/
	
	var arr1 = [3]int{1,2,3};
	fmt.Println(arr1);

	var arr2 = [...]int{4,5,6,7};
	fmt.Println(arr2);

	/*
		2. With the := sign
		
			array_name := [length]data_type{values}
			or
			array_name := [...]data_type{values}
	*/	

	arr3 := [2]int{11,22};
	fmt.Println(arr3);

	arr4 := [...]int{33,44,55};
	fmt.Println(arr4);

	var cars = [4]string{"Car-1", "Car-2", "Car-3", "Car-4"};
	fmt.Println(cars[1]);
	fmt.Println(cars[3]);

	cars[1] = "Car-22";
	fmt.Println(cars[1]);

	/*
		1. //not initialized
		2. //partially initialized
		3. //fully initialized
	*/

	arr5 := [3]int{};  //not initialized
	arr6 := [4]int{1,2}  //partially initialized
	arr7 := [5]int{1,2,3,4,5} // fully initialized

	fmt.Println(arr5);
	fmt.Println(arr6);
	fmt.Println(arr7);

	arr8 := [5]int{1:10,4:40}
	fmt.Println(arr8);

	fmt.Println(len(arr8));
	
}