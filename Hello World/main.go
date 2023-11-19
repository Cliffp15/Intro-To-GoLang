package main

import (
	"fmt"
	"time"
)

func printFunc(name string) {
	for i := 1; i < 4; i++ {
		fmt.Println(name+": ", i)
	}
}

func main() {
	var str string = "Hello World!"
	fmt.Println(str)

	fmt.Println("Showing off GoRoutines")
	fmt.Println()
	fmt.Println()
	// running go routine using existing function
	go printFunc("GoRoutine 1")
	// running function in main thread
	printFunc("Main thread")
	// running anonymous go routine
	go func() {
		fmt.Println("Anonymous function GoRoutine")
	}()
	// Sleep is needed so main process doesn't exit until all is done
	time.Sleep(time.Second * 5)
	fmt.Println("Done")
}
