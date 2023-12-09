package vectors

import "fmt"

func Vectors(){
	var myVector [6]int
	fmt.Println(myVector)

	myarr := [3]int{1,2,3}
	fmt.Println(myarr)

	myarr[2] = 5
	fmt.Println(myarr)


	var silice1 []int
	fmt.Println(silice1)
	silice1 = append(silice1, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(silice1)
	silice2 := silice1[1:3]
	fmt.Println(silice2)
	fmt.Println(&silice1[2])
	fmt.Println(&silice2[1])
	silice2[0] = 100
	fmt.Println(silice1)
	fmt.Println(silice2)

}