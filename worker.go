package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		id += 1
		fmt.Println(id)
		results <- j

	}
}
func main() {

	jobs := make(chan int, 1000)
	results := make(chan int, 1000)

	for w := 1; w <= 10; w++ {
		go worker(0, jobs, results)
	}

	for j := 1; j <= 1000; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 1000; a++ {
		<-results

	}

}
