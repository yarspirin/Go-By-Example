package main

import (
	"fmt"
	"time"
)

func main() {
	nums := make(chan int, 5)

	for i := 0; i < 5; i++ {
		nums <- i
	}

	close(nums)

	limiter := time.Tick(200 * time.Millisecond)

	for num := range nums {
		<-limiter
		fmt.Println(num, time.Now())
	}

	fmt.Println()

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for tick := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- tick
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println(req, time.Now())
	}
}
