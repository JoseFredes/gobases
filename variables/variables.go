package variables

import "fmt"

// Types of variables

// default values
// mumeric 0
// string ""
// boolean false
// pointers nil
// interfaces nil
// functions nil
	

func Variables() {

	var myInVar int
	var myStrVar string
	var myBoolVar bool
	myInVar = 10
	myStrVar = "Hello"
	myBoolVar = true

	fmt.Println("myInVar: ", myInVar)
	fmt.Println("myStrVar: ", myStrVar)
	fmt.Println("myBoolVar: ", myBoolVar)

	var uintVar uint
	uintVar = 10

	fmt.Println("uintVar (only possitive): ", uintVar)

	fmt.Println("My variable address", &myInVar)

	// Short declaration
	myShortVar := 10
	fmt.Println("myShortVar: ", myShortVar)


	const myConstVar int = 10
	fmt.Println("myConstVar: ", myConstVar)

	const myIntVar = "a10"
	fmt.Println("myIntVar: ", myIntVar)

	var my8BitInt int8
	fmt.Printf("my8BitInt: %d\n", my8BitInt)
}