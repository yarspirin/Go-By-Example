package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 2)

	<-timer.C // waits for two seconds (equivalent to time.Sleep(...))

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer expired.")
	}()

	stop := timer2.Stop()

	if stop {
		fmt.Println("Timer stopped.")
	}
}
