package main

import (
	"fmt"
	"time"
	"sync"
	"context"

	"golang.org/x/sync/semaphore"
)


func SendPromoEmail(uid int) {
	//simulate email sending
	fmt.Printf("sending email to user %d\n", uid)
	time.Sleep(100*time.Millisecond)
}


func main() {
	totalEmails := 10000
	maxConcurrent := 50
	
	ctx := context.Background()
	sem := semaphore.NewWeighted(int64(maxConcurrent))
	var wg sync.WaitGroup

	for i := 1; i <= totalEmails; i++ {
		wg.Add(1)

		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Println("failed to acquire semaphore: %v\n", err)
			wg.Done()
			continue
		}

		go func(uid int) {
			defer wg.Done()
			defer sem.Release(1)

			SendPromoEmail(uid)
		}(i)
	}

	wg.Wait()
}


