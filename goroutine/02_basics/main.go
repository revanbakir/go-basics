package main


import (
"fmt"
"sync"
)


func task(name string, wg *sync.WaitGroup) {
	defer wg.Done()   // 
	fmt.Println(name, "completed")
}



func main() {

	var wg sync.WaitGroup // içinde değeri 0 olan sayaç tutan waitgroup oluşturulur
	wg.Add(2)  // sayaç 2. 2 görevin bitmesini beklicem

	go task("task 1", &wg)  // wg değişkeninin kendisinin adresini ver
	go task("task 2", &wg)

	wg.Wait()  // main, sayaç 0 olana kadar bloklanır
	fmt.Println("all tasks completed")
}
	
