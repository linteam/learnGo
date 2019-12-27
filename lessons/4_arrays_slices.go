package lessons

import(
	"fmt"
)

func LearnArrays() {
	//Declare and Assign
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	//Shortway
	personArr := [2]string{"Ahmed","Ali"}

	fmt.Println(personArr)
	fmt.Println(fruitArr)

}