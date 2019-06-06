package main

import "fmt"

func main() {
	jobs := make(chan int, 100)
	done := make(chan bool)

	go func() {
		for {
			job, isDone := <-jobs

			if isDone {
				fmt.Println("All jobs have been finished.")
				done <- true
				break
			} else {
				fmt.Printf("Job %d has been finished.\n", job)
			}
		}
	}()

	for job := 1; job < 10; job++ {
		jobs <- job
		fmt.Printf("Job %d is going to be running\n", job)
	}

	close(jobs)
}
