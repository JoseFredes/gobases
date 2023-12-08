package operators

import "fmt"


func Operators(){
	yearsOld := 10

	fmt.Println(yearsOld > 15)
	fmt.Println(yearsOld < 15)
	fmt.Println(yearsOld == 15)
	fmt.Println(yearsOld != 15)
	fmt.Println(yearsOld >= 15)
	fmt.Println(yearsOld <= 15)
	 
	// Logical operators
	fmt.Println(yearsOld > 15 && yearsOld < 20)
	fmt.Println(yearsOld > 15 || yearsOld < 20)
	fmt.Println(!(yearsOld > 15 && yearsOld < 20))

	// not operator
	fmt.Println(!true)
	fmt.Println(!false)

	// Condiional operators
	fmt.Println("Conditional operators")
	if yearsOld > 15 {
		fmt.Println("You can drive")
	} else {
		fmt.Println("You can't drive")
	}

	// Ternary operator
	fmt.Println("Ternary operator")
	canDrive := yearsOld > 15
	fmt.Println(canDrive)

	if canDrive {
		fmt.Println("You can drive")
	} else if yearsOld == 15 {
		fmt.Println("You can't drive")
	} else {
		fmt.Println("You can't drive")
	}

	if value := false; value {
		fmt.Println("Value is true")
	} else {
		fmt.Println("Value is false")
	}


}