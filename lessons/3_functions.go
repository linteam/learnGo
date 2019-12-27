package lessons
import(
	"fmt"
)
func greeting(name string) string {
	return "Selam " + name
}
func getSum(num1, num2 int) int{
	return num1+num2
}
func functionsLearn(){
	fmt.Println(greeting("Ali"))
	fmt.Println(getSum(2,9))
}