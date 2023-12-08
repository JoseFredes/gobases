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

	// Switch
	fmt.Println("Switch")
	switch yearsOld {
	case 10:
		fmt.Println("You are 10 years old")
	case 15:
		fmt.Println("You are 15 years old")
	default:
		fmt.Println("You are not 10 or 15 years old")
	}

	// Switch with no condition
	fmt.Println("Switch with no condition")
	switch {
	case yearsOld < 10:
		fmt.Println("You are less than 10 years old")
	case yearsOld > 10 && yearsOld < 15:
		fmt.Println("You are between 10 and 15 years old")
	default:
		fmt.Println("You are not 10 or 15 years old")
	}


}