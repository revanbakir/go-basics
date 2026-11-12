package main


import (
"fmt"
"sync"
)


func task(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(name, "completed")
}



func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go task("task 1", &wg)
	go task("task 2", &wg)

	wg.Wait()
	fmt.Println("all tasks completed")
}
	
