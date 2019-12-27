package lessons

import(
	"fmt"
)

//LearnArrays ...
func LearnArrays() {
	//Declare and Assign
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	//Shortway
	//personArr := [2]string{"Ahmed","Ali"}
	students := []string{"Ahmed","Ali","Harun","Hakki"}

	fmt.Println("All members of students array:",students)
	fmt.Println("First 2 members of students:",students[0:2])
	fmt.Println("Length of Fruit array:",len(fruitArr))



}