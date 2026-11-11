package main

import (
"fmt"
"time"
)


func say(s string) {
	for _ = range 5 {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}


func main() {
	
	go say("hello")
	say("world")	
	
}
		
