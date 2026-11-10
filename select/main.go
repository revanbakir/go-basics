package main

import (
"fmt"
"time"
)

func main() {

	k1 := make(chan string)
	k2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		k1 <- "video"
	}()

	go func () {
		time.Sleep(time.Second*3)
		k2 <- "ses"
	}()


	for i := 0; i < 2; i++ {
		select {
		case mesaj1 := <- k1:
			fmt.Println("mesaj 1 : ", mesaj1)
		case mesaj2 := <- k2:
			fmt.Println("mesaj 2 : ", mesaj2)
		}
	}
} 
