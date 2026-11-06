package main 

import (
"fmt"
"time"
)


// goroutine ile aynı runtime içinde multiple thread oluşturabiliriz (concurrency)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func yazdır(yazı string) {
	fmt.Println(yazı)
}

func main() {

	f("direct")  // senkron (doğrudan)

	go f("goroutine")


	go func(msg string) {
		fmt.Println(msg)
	}("going")


	fmt.Println("hellworld")
	go yazdır("eşzamanlı işlem")



	time.Sleep(time.Second)
	fmt.Println("done")

}


