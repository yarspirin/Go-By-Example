package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started working on job %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished working on job %d\n", id, job)
		results <- job * job
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for job := 1; job < 15; job++ {
		jobs <- job
	}

	close(jobs)

	for res := 1; res < 15; res++ {
		fmt.Println(<-results)
	}
}
