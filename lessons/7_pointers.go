package lessons

import (
	"fmt"
)

//LearnPointers ...
func LearnPointers()  {
	a := 5
	b := &a
	fmt.Println(a, b)	
	fmt.Printf("a'nin tipi %T\n",a)
	fmt.Printf("b'nin tipi %T\n",b)

	// Use * to read values from address
	fmt.Println("b adresindeki deger: ",*b)
	fmt.Println("a degerinin adresindeki deger = a: ",*&a)

	*b = 10
	fmt.Println("*b atamasi ile a degeri degisti mi? : ",a)


}