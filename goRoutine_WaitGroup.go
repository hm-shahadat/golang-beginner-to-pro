package main

import (
	"fmt"
	"sync"
)

func goRoutine(valu int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("This is the value: ", valu)

}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go goRoutine(i, &wg)
	}
	wg.Wait()
}
