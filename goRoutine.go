package main

import (
	"fmt"
	"time"
)

// func goRoutine(valu int) {
// 	fmt.Println("This is the value: ", valu)

// }

func main() {
	for i := 0; i <= 10; i++ {
		// 	go goRoutine(i)
		go func(i int) {
			fmt.Println("This is the value: ", i)
		}(i)
	}

	time.Sleep(time.Second * 1)
}
