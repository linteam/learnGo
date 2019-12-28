package lessons

import (
	"fmt"
	"strconv"
)

//Person ...
// type Person struct{
// 	firstName string
// 	lastName string
// 	city string
// 	gender string
// 	age int
// 	childCount int
// }
type Person struct{
	firstName, lastName, city, gender string
	age, childCount int
}

//Greeting method (value receiver) 
//Person struct'ina greet metodu ekledim.
func (p Person) greet() string{
	message := "Hello, my name is " + p.firstName + " " + p.lastName
	message += " and I am " + strconv.Itoa(p.age)
	return message
}

//hasbirthday (pointer receiver) degisiklik yapacagiz
func (p *Person) hasbirthday() {
	p.age++
}

func (p *Person) getMarried(childCount int)  {
	p.childCount += childCount;
}

//Structs ...
func Structs(){
	//Init (Ustteki yontem tercih edilebilir)
	//ali := Person{firstName:"Ali",lastName:"Celik",city:"Kars",gender:"m",age:34}
	veli := Person{"Ali","Celik","Kars","m",34,0}

	fmt.Println(veli)
	fmt.Println(veli.lastName)

	//Struct'a 2 tip fonksiyon ekleyebiliyoruz, Value ve reference tipli
	//Valut tip
	veli.hasbirthday() //dogum gununu bir arttirdi
	fmt.Println(veli.greet())

	veli.getMarried(3)
	fmt.Println("Cocuk sayisi: ",veli.childCount)

}