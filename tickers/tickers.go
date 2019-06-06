package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		for tick := range ticker.C {
			fmt.Println(tick)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
}
