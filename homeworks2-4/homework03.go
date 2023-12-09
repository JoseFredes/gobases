package homeworks2_4

import "fmt"



func Homework03(){
	var number int
	numbers := []int{}

	for {
		fmt.Println("Enter a number: ")
		fmt.Scanln(&number)

		if number == 0 {
			break
		}

		numbers = append(numbers, number)
	}

	fmt.Println("The numbers are: ", numbers)
}