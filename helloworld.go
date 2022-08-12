package main

import "fmt"

/* If variable defined outside function and not used,
It will give warning but will execute the program */
var age = 29

/* If constant defined outside function and not used,
It will not give eror or warning and will execute the program
as usual */
const AGE = 29

/* We can't declare a variable outside a function with := */
//book := "The Monk who sold his Ferari"

func main() {
	const NAME = "Ganesh"
	fmt.Println(NAME)

	/* If variable defined inside a function a not used,
	It will give a error message and will not execute the program */
	var address = "Jamshedpur"
	address = "Nagpur"
	fmt.Println(address)

	/* If a constant is defined inside a function and not used,
	It will not give any error but give warning and will execute the program*/
	const ADDRESS = "Bokaro"

	/*Variable Declaration Without Initial Value.
	Defualt value for string is "" empty string, for int it is 0 and
	for bool it is false. */
	var a string
	var b int
	var c bool
	fmt.Print("a: ", a, " ", "b: ", b, " ", "c: ", c, "\n")

	// Multiple Variable Declaration stricting to use same data type
	var aa, bb, cc, dd int = 1, 3, 5, 7
	fmt.Println(aa, bb, cc, dd)

	// Multiple Variable Declaration with multi data type
	var ab, bc = 6, "Hello"
	cd, de := 7, "World!"
	fmt.Println(ab, bc, cd, de)

	// Variable Declaration in a Block
	var (
		x int
		y int    = 12
		z string = "Hello"
	)
	fmt.Println(x, y, z)

	/* Multiple constant declaration*/
	const (
		A int = 30
		B     = 3.41
		C     = "Hi"
	)
	fmt.Print("A: ", A, " ", "B: ", B, " ", "C: ", C, "\n")

	/* Camel Case: multi word variable declaration,
	Each word, except the first, starts with a capital letter */

	var myVariableName = "John"
	fmt.Println(myVariableName)

	/* Pascal case: multi word variable declaration,
	Each word starts with a capital letter */

	var MyVariableName = "John"
	fmt.Println(MyVariableName)

	/* Snake Case: multi word variable declaration,
	Each word is separated by an underscore character*/

	var my_variable_name = "John"
	fmt.Println(my_variable_name)

	/* Output Functions: Go has three functions to output value
	Print(): It prints its arguments with their default format

	Println(): The Println() function is similar to Print()
	with the difference that a whitespace is added between the arguments,
	and a newline is added at the end.

	Printf(): The Printf() function first formats its
	argument based on the given formatting verb and then prints them */

	// The Print() function prints its arguments with their default format.

	var i, j string = "Hello", "World"

	fmt.Print(i)
	fmt.Print(j)

	// If we want to print the arguments in new lines, we need to use \n.

	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	//It is also possible to only use one Print() for printing multiple variables.

	fmt.Print(i, "\n", j, "\n")

	// If we want to add a space between string arguments, we need to use " ":

	fmt.Print(i, " ", j, "\n")

	/* Print() inserts a space between the arguments
	if neither are strings by default */

	var k, l int = 2, 4
	fmt.Print(k, l)

	/* The Println() function is similar to Print() with the difference
	that a whitespace is added between the arguments, and a newline
	 is added at the end. */

	var m, n string = "Hello", "World"
	fmt.Println(m, n)

	/* The Printf() function first formats its argument based on the
	given formatting verb and then prints them.
	%v is used to print the value of the arguments
	%T is used to print the type of the arguments
	%% Prints the % sign
	%#v Prints the value in Go-syntax format */

	var o string = "Hello"
	var f int = 15
	var r = false
	var s = true
	fmt.Printf("o has a value: %q and type: %T\n", o, o)
	fmt.Printf("f has a value: %v and type: %T and it has %v%%\n", f, f, f)
	fmt.Printf("%t\n", r) // boolean formatting verbs
	fmt.Printf("%t\n", s) // boolean formatting verbs

	/* Arrays are used to store multiple values of the same type in a single variable,
	instead of declaring separate variables for each value.
	In Go, there are two ways to declare an array:
	1) With the var keyword
	2) With the := sign
	*/

	// Below is example of array declaration with var keyword.
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	var arr2 = [5]string{"Mahindra", "Tata", "Ferari", "Skoda", "BMW"}
	fmt.Println(arr2)

	// Access elements of array
	fmt.Println(arr2[0])

	// Changing element of an array
	arr2[2] = "Bugati"
	fmt.Println(arr2)

	var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr1)

	/* If an array or one of its elements has not been initialized in the code,
	it is assigned the default value of its type */

	// In below array elements have not been initialized.
	var arr3 = [5]int{}
	fmt.Println(arr3)

	// In below array elements have been partially initialized
	var arr4 = [5]int{1, 2, 3}
	fmt.Println(arr4)

	// In below array elements have been fully initialized.
	var arr5 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr5)

	// Initialize only at specific index

	var arr6 = [5]int{1: 3, 2: 5}
	fmt.Println(arr6)

	// Find the length of an array.
	fmt.Println(len(arr6))

	/*
		Go Slices
		Slices are similar to arrays, but are more powerful and flexible.
		Like arrays, slices are also used to store multiple values of the same type in a single variable.
		However, unlike arrays, the length of a slice can grow and shrink as you see fit.
		In Go, there are several ways to create a slice:
		Using the []datatype{values} format
		Create a slice from an array
		Using the make() function
	*/

	// Create a Slice With []datatype{values}

	var mySlice = []int{1, 2, 3, 4, 5}
	mySlice = append(mySlice, 50, 60)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	myMessage() // myMessage function is called.
	myPara("Ganesh")
	multiPara("Ganesh", 29)
	fmt.Println(myFunction(5, 6))
	fmt.Println(myNamed(10, 30))

	// Store the return value of function

	var total int = myNamed(20, 10)
	fmt.Println(total)

	fmt.Println(multiReturn(5, "Hello"))

	// We can store multi return values in different variables

	q, w := multiReturn(6, "Namaste")
	fmt.Println(q, w)

	/* If we (for some reason) do not want to use some of the returned values,
	we can add an underscore (_), to omit this value. */

	//_, t := multiReturn(8, "Bon Jour")
	//fmt.Println(t)
	//
	//testCount(1)

	//factorial_recursion(4)

	/* */

	/* Maps are used to store data values in key:value pairs.
	A map is an unordered and changeable collection that does not allow duplicates. 
	*/
	var testmap = map[string]string{"fname": "Ganesh", "lname": "Ram", "City": "Jamshedpur", "Job": "Software engineer"}
	testmap1 := map[string]int{"Oslo": 10, "Bergen": 20, "Delhi": 30, "Mumbai": 40}
	fmt.Printf("testmap\t%v\n", testmap)
	fmt.Printf("testmap1\t%v\n", testmap1)

	// Creating Maps Using Using make()Function:
	var 

}

/* A function is a block of statements that can be used repeatedly in the programme.
A function can not be run automatically when the page is loaded, A function will be executed
when the function is called. */

func myMessage() {
	fmt.Println("I just got executed......")
}

// passing parameters in the function.

func myPara(fname string) {
	fmt.Println("Hello", fname, "how are you doing today...")
}

// Multi parameter function

func multiPara(fname string, age int) {
	fmt.Println("Hello", fname, "you are ", age, "year old....")
}

// Go function returns

func myFunction(x int, y int) int {
	return x + y
}

// Named return values

func myNamed(a int, b int) (result int) {
	result = a * b
	return //result
}

func multiReturn(x int, y string) (result int, txt string) {
	result = x + x
	txt = y + " Ganesh"
	return
}

// A function is recursive if it calls itself and reaches a stop condition.

func testCount(z int) int {
	if z == 11 {
		return 0
	}
	fmt.Println(z)
	return testCount(z + 1)
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
		fmt.Println(y)
	} else {
		y = 1
		fmt.Println(y)
	}
	return
}

/*
A struct (short for structure) is used to create a collection of members
of different data types, into a single variable.


*/

type Person struct {
	name   string
	age    int
	job    string
	salary int
}
