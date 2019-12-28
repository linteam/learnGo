package lessons

import "fmt"

//Loops ...
func Loops()  {
	for i := 0; i < 5; i++ {
		fmt.Printf("Number: %d\t",i) //\t ile tab olur
	}	
}

//Maps for key value dictionary
func Maps(){
	// //Define Map
	// emails := make(map[string]string)
	// //Assign kv
	// emails["Berat"] = "berat@gmail.com"	

	//Declare and add kv
	emails := map[string]string{"Berat":"berat@gmail.com"}
	emails["Apti"] = "karademir@tubitak.gov.tr"
	emails["Ercan"] = "sahaprojesi@yahoo.com"
	
	fmt.Println("Map Length:",len(emails))
	fmt.Println(emails["Apti"])

	//Delete
	delete(emails, "Apti")
	fmt.Println("Apti deleted...",emails)

	for k,v:=range emails{
		fmt.Printf("key: %s - value: %s\n",k,v)
	}

	for k := range emails{
		fmt.Printf("Key Only: %s\n",k)
	}

	for _,v := range emails{
		fmt.Printf("Value Only: %s\n",v)
	}
}

//Ranges ...
func Ranges(){
	ids := []int{33,54,54,34,12,43,54}
	//Loop through ids
	for i, id:= range ids{
		fmt.Printf("%d - ID: %d\n",i, id)
	}

	//Not using index
	sum := 0
	for _, id:= range ids{
		sum += id
	}
	fmt.Println("Sum",sum)
}