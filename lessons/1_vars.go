package lessons //main fonksiyonun cagrilmasi icin paket tanimlanmali
import "fmt"

var name string = "Hakan"

//Bu otomatik olarak cagrilir.
func learnVariables() {
	/* MAIN TYPES*/
	//Go'da bir degisken tanimlandiginda onu kullanman gerekli
	second := "Tayyip"	//Bu tanimlama seklini fonksiyon disinda yapamiyoruz
	var surname = "Pekdemir"
	zName, zSurname := "Hakki","Ocal"
	//var isCrazy bool = false
	//var age int = 34        //int8 int16 int32 int64
	//var experience uint = 5 //uint8 uint16 uint32 uint64
	//var data byte = 0       //alias for uint8
	//var data2 rune = 0      //alias for int32
	//var iq float32 = 45.43  //float32 float64
	size := 1.3 //default float64 gelir
	//complex64 complex128 for really large numbers

	const dataCanRead = "Pasaoglu" //const degeri kullanmayinca patlamadi :)
	fmt.Println("Selam" , name , second, surname) //, bi bosluk birakir
	fmt.Println(zName,zSurname)
	fmt.Printf("%T\n", size) //T type alinir, sonuc: int
}
