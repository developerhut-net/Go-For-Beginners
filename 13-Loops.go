package main

import "fmt"


func main(){

/*

for statement-1; statement-2; statement-3 {
	//code block
} 

statement-1 : initialize
statement-2 : condition 
statement-3 : update 

*/

for i:=0; i<5; i++ {

	if i==3 {
		continue
	}

	if i==4 {
		break
	}

	fmt.Println(i);
}

arr1 := [2]int{1,2}
arr2 := [3]int{1,2,3}

for i:=0; i < len(arr1); i++ {

	for j:=0; j < len(arr2); j++ {

		fmt.Printf("%v * %v = %v\n", arr1[i], arr2[j], arr1[i]*arr2[j])
	}

}

// Ranges

/*
	for index, value := range array|slice|map {
		// code block
	}
*/

country := [3]string{"UAE", "England", "South Africa"};

for index, value := range country {
	fmt.Printf("%v\t%v\n", index, value);
}

for _, value := range country {
	fmt.Printf("%v\n", value);
}

for i, _ := range country {
	fmt.Printf("%v\n", i);
}

}