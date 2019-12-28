package lessons

import (
	"fmt"
)

//adder fonskiyonu, int alip int donen bir fonksiyon doner
func adder() func(int) int{
	sum:=0
	return func(x int) int {
		sum += x
		return sum
	}
}
//Closures ...
func Closures()  {
	sum := adder() //sum fonksiyonunu aldim
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}