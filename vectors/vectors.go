package vectors

import "fmt"

func Vectors(){
	var myVector [6]int
	fmt.Println(myVector)

	myarr := [3]int{1,2,3}
	fmt.Println(myarr)

	myarr[2] = 5
	fmt.Println(myarr)
}