package homeworks2_4

import "fmt"

func Homework02(){
	array := [5]int{4, 2, 5, 6, 7}
    
    for idx, value := range array {
		value += 20
		array[idx] = value
	}
 
    fmt.Println("Los valores del array son: ", array)
}