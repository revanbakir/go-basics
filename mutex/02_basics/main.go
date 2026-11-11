package main

import (
"fmt"
"sync" // mutex ve waitgroup için
)


// Birden fazla goroutine in aynı map'i güvenli şekilde güncellenmesi

type Container struct {   
	mu sync.Mutex
	counters map[string]int   // isim -> sayaç değerini tutan map
}


func (c *Container) inc(name string) {  // *Container(pointer) çünkü her çağrı aynı structtaki aynı mutexi kullanır 

	c.mu.Lock()
	defer c.mu.Unlock()  
	c.counters[name]++
}


func main() {

	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}


	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {
		for range n { 
			c.inc(name)
		}
	}


	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
        	doIncrement("a", 10000)
    	})

	wg.Go(func() {
        	doIncrement("b", 10000)
    	})

	wg.Wait()
	fmt.Println(c.counters)
}
