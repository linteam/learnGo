package lessons

import (
	"fmt" 
	"math"
	"./strutil"
)
func packagesLearn(){
	fmt.Println(math.Abs(-2)) //2	
	fmt.Println(math.Ceil(2.7))	//3	
	fmt.Println(math.Floor(2.7)) //2
	fmt.Println(math.Sqrt(16)) //4

	fmt.Println(strutil.Reverse("bu metin tersine donecek"))
}