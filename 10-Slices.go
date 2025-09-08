package main

import "fmt"


func main(){

	/*
		1. []datatype{values} format
		2. Create a slice from an array
		3. make() function	
	*/

	//slice_name := []datatype{values}
	slice1 := []int{}

	slice2 := []int{1,2,3}

	//len() & cap()

	fmt.Println(len(slice1));
	fmt.Println(cap(slice1));
	fmt.Println(slice1);


	fmt.Println(len(slice2));
	fmt.Println(cap(slice2));
	fmt.Println(slice2);

	
	//slice_name := array[start:end]
	var array = [6]int{1,2,3,4,5,6}
	slice3 := array[2:4]

	fmt.Println(len(slice3));
	fmt.Println(cap(slice3));
	fmt.Println(slice3);


	//make()
	//slice_name = make([]type, lenght, capacity)
	slice4 := make([]int, 5, 10);

	fmt.Println(len(slice4));
	fmt.Println(cap(slice4));
	fmt.Println(slice4);

	slice5 := make([]int, 5);
	fmt.Println(len(slice5));
	fmt.Println(cap(slice5));
	fmt.Println(slice5);

	slice6 := []int{10,20,30,40};
	fmt.Println(slice6[1]);
	fmt.Println(slice6[3]);

	slice6[0] = 100;
	fmt.Println(slice6[0]);

	//slice_name = append(slice_name, ele1, ele2, ....)
	slice7 := []int{1,2};
	fmt.Printf("slice : %v\n", slice7);
	fmt.Printf("length : %v\n", len(slice7));
	fmt.Printf("capacity : %v\n", cap(slice7));

	slice7 = append(slice7, 3, 4);
	fmt.Printf("slice : %v\n", slice7);
	fmt.Printf("length : %v\n", len(slice7));
	fmt.Printf("capacity : %v\n", cap(slice7));


	//slice_name := append(slice1, slice2...)
	slice8 := []int{1,2};
	slice9 := []int{3,4};

	var slice10 = append(slice8, slice9...);
	fmt.Printf("slice : %v\n", slice10);
	fmt.Printf("length : %v\n", len(slice10));
	fmt.Printf("capacity : %v\n", cap(slice10));


	//length
	arr1 := [6]int{1,2,3,4,5,6};
	slice11 := arr1[1:5];
	fmt.Printf("slice11 : %v\n", slice11);
	fmt.Printf("length : %v\n", len(slice11));
	fmt.Printf("capacity : %v\n", cap(slice11));

	slice11 = arr1[1:3]
	fmt.Printf("slice11 : %v\n", slice11);
	fmt.Printf("length : %v\n", len(slice11));
	fmt.Printf("capacity : %v\n", cap(slice11));

	slice11 = append(slice11, 10, 11, 12, 13);
	fmt.Printf("slice11 : %v\n", slice11);
	fmt.Printf("length : %v\n", len(slice11));
	fmt.Printf("capacity : %v\n", cap(slice11));

}