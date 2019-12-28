package lessons
import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello")
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>About</h1>")
}

//WebSample ...
func WebSample(){
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server starting at port 3000...")
	http.ListenAndServe(":3000",nil) //null nil go dilinde
}