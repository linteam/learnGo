package lessons

import "fmt"

//LearnConditionals ...
func LearnConditionals()  {
	x := 15
	y := 10
	if x < y { //<=, >= kullanilabilir.
		//fmt.Println("") %d yazma sekli Println'de calismiyor
		fmt.Printf("%d is less than %d\n",x,y)
	}else{
		fmt.Printf("%d is greater than %d\n",x,y)
	}

	color := "red"
	if color == "red"{
		fmt.Println("stop the car")
	}
	//else if, else ...
	color = "blue"
	switch color {
	case "red":
		fmt.Println("stop the car")
	case "green":
		fmt.Println("go go go")
	default:
		fmt.Println("renk körüsün")		
	}
}