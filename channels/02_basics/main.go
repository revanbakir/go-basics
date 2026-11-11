package main

import (
"fmt"
"time"
)


func main() {

	go func() {
		time.Sleep(time.Second * 2)
			fmt.Println("ilk yazı")
	}()
	fmt.Println("ikinci yazı")
	time.Sleep(time.Second * 1)	
}
