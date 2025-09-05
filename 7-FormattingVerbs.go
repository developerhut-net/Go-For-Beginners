package main

import "fmt"


func main(){

	/*
		1. General Formatting Verbs
		2. Integer Formatting Verbs
		3. String Formatting Verbs
		4. Boolean Formatting Verbs
		5. Float Formatting Verbs	
	*/

/*
	1. General Formatting Verbs
	Verb	Description
	----	------------
	%v		Prints the value in the default format
	%#v		Prints the value in Go-syntax format
	%T		Prints the type of the value
	%%		Prints the % sign
*/

	var i = 15.5;
	var j string = "Hello";

	fmt.Printf("%v\n", i);
	fmt.Printf("%#v\n", i);
	fmt.Printf("%T\n", i);
	fmt.Printf("%v%%\n", i);

	fmt.Printf("%v\n", j);
	fmt.Printf("%#v\n", j);
	fmt.Printf("%T\n", j);

/*
	2. Integer Formatting Verbs
	Verb	Description
	----	------------
	%d		Base 10
	%+d		Base 10 and always show sign
	%4d		Pad with spaces (width 4, right justified)
	%-4d	Pad with spaces (width 4, left justified)
	%04d	Pad with zeroes (width 4)
*/

	var x = 10;
	fmt.Printf("%d\n", x);
	fmt.Printf("%+d\n", x);
	fmt.Printf("%5d\n", x);
	fmt.Printf("%-5d\n", x);
	fmt.Printf("%05d\n", x);


/*
	3. String Formatting Verbs
	Verb	Description
	----	------------
	%s		Prints the value as plain string
	%q		Prints the value as a double-quoted string
	%8s		Prints the value as plain string (width 8, right justified)
	%-8s	Prints the value as plain string (width 8, left justified)
	%x		Prints the value as hex dump of byte values
*/

	var z string = "John";

	fmt.Printf("%s\n", z);
	fmt.Printf("%q\n", z);
	fmt.Printf("%8s\n", z);
	fmt.Printf("%-8s\n", z);
	fmt.Printf("%x\n", z);


/*
	4. Boolean Formatting Verbs
	Verb	Description
	----	------------
	%t		Value of the boolean operator in true or false 
				format (same as using %v)
*/

	var a, b = true, false;

	fmt.Printf("%t\n", a);
	fmt.Printf("%t\n", b);

/*
	5. Float Formatting Verbs
	Verb	Description
	----	------------
	%e		Scientific notation with 'e' as exponent
	%f		Decimal point, no exponent
	%.2f	Default width, precision 2
	%6.2f	Width 6, precision 2
*/

	var c = 4.123;

	fmt.Printf("%e\n", c);
	fmt.Printf("%f\n", c);
	fmt.Printf("%.2f\n", c);
	fmt.Printf("%6.3f\n", c);


}