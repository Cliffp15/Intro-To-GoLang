package main 


import(
	"fmt"
	"runtime"
)

func main(){
	var str string ="Hello World"

	fmt.Println(str)

	runtime.Breakpoint()
	var str1 string 
	str1 = "Goodbye World"
	fmt.Println(str1)


}