package main

import (
"fmt"
"time"
)

func main() {

	c := make(chan bool)	// bool türünde bir kanal oluştur
	go hello(c)	// hello yu çağıran ve kanala veri gönderen bir goroutine oluştur
	fmt.Println("main fonksiyon")
	<-c		// goroutine nin kanala veri göndermesini bekle
	fmt.Println("main sona eriyor")
	

}


func hello(c chan bool) {
	fmt.Println("merhaba dünya")
	time.Sleep(2*time.Second) 
	fmt.Println("tekrar merhaba")
	c <-true
}
