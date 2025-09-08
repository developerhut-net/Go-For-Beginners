package main

import "fmt"


func main(){

	/*
		bool
		numeric 
		string
	*/

	var b1 bool = true;
	var b2 = true;
	var b3 bool;
	b4 := true;

	fmt.Println(b1);
	fmt.Println(b2);
	fmt.Println(b3);
	fmt.Println(b4);

	/*
		1. Signed integers 
		2. Unsigned integers
	*/

	//var x int = 100;
	//var y int = -100;

	/*
	
		Type	Size								Range
		----	----								-----
		int		Platform dependent :				-2147483648 to 2147483647 in 32 bit systems and
				32 bits in 32 bit systems and		-9223372036854775808 to 9223372036854775807 in 64 bit systems
				64 bits in 64 bit systems	
		int8	8 bits/1 byte						-128 to 127
		int16	16 bits/2 byte						-32768 to 32767
		int32	32 bits/4 byte						-2147483648 to 2147483647
		int64	64 bits/8 byte						-9223372036854775808 to 9223372036854775807
	
	*/	

	//var a uint = 100;
	//var b uint = 300;

	/*
	
		Type	Size								Range
		----	----								-----
		uint	Depends on platform:				0 to 4294967295 in 32 bit systems and
				32 bits in 32 bit systems and		0 to 18446744073709551615 in 64 bit systems
				64 bit in 64 bit systems			
		uint8	8 bits/1 byte						0 to 255
		uint16	16 bits/2 byte						0 to 65535
		uint32	32 bits/4 byte						0 to 4294967295
		uint64	64 bits/8 byte						0 to 18446744073709551615

	*/
	//c := 7;
	//var d = 8;
	// default data type - int

	/*
		Type		Size		Range
		----		----		-----
		float32		32 bits		-3.4e+38 to 3.4e+38
		float64		64 bits		-1.7e+308 to +1.7e+308
	*/

	//var e float32 = 127.55;
	//var f float64 = 1.6e+209;

	// default data type - float64

	var text1 string = "Hello";
	var text2 string;
	text3 := "World";

	fmt.Printf("Type: %T, value: %v\n", text1, text1);
	fmt.Printf("Type: %T, value: %v\n", text2, text2);
	fmt.Printf("Type: %T, value: %v\n", text3, text3);

}