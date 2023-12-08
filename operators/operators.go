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

}