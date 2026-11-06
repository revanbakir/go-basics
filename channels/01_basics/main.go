package main 

import (
"fmt"
"time"
)


func main() {

	messages := make(chan string)  // yeni bir kanal

	go func() { messages <- "ping"}()  // kanala değer gönderme
	
	
	msg := <-messages  // kanaldan mesajı alır ekrana yazdırır
	fmt.Println(msg)



	k := make(chan bool)
	

	go func() {
		time.Sleep(time.Second*5)
		k <- true
	}()
			
	message := <-k
	fmt.Println(message)


	c := make(chan bool, 2)  // 2 adet bool değer taşıyan channel
	
	go func() {   // senkron thread
		time.Sleep(time.Second * 5)
		c <- true
	
		time.Sleep(time.Second * 2)
		c <- false
	}()

	fmt.Println(<-c)


}
