package main

import (
	"fmt"
	"sync"
)


func main() {
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
			fmt.Println(counter)
		}()
	}

	wg.Wait()
	fmt.Println("counter: ", counter)

// ***********************************************************************
// fixing a race condition using a mutex

	var mu sync.Mutex
	counter2 := 0

	for m := 0; m < 100; m++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			counter2++
			fmt.Println(counter2)
		}()
	}
	wg.Wait()
	fmt.Println("counter2: ", counter2)

}
