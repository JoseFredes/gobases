package forloop

import "fmt"

func Forloop(){
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		break
	}
	
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
		break
	}
	
	for {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
		break
	}
	
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		break
	}

	mySlice := []int{1,2,3,4,5,6,7,8,9}
	for i, value := range mySlice {
		fmt.Println(i, value)
	}

}