package main


import (

	"fmt"
	"sync"
	"time"
)


func main() {
	totalEmails := 10000
	maxConcurrent := 50  // it is a cap on how many goroutines run at once 

	sem := make(chan struct{}, maxConcurrent)  // max 50 emails concurrently struct{} is an empty type, zero space
	var wg sync.WaitGroup

	for i := 1; i <= totalEmails; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(uid int) {
			defer wg.Done()
			defer func() { <-sem }()
			SendPromoEmail(uid)
		}(i)
	}

	wg.Wait()  	
}


func SendPromoEmail(uid int) {
	// Simulate email sending
	fmt.Printf("sending email to user %d\n", uid)
	time.Sleep(100*time.Millisecond)	
}
