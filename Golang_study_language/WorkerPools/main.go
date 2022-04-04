package main

import (
	"fmt"
	"time"
)

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= numWorkers; w++ {
		fmt.Println("w ---> ", w)
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		fmt.Println(" Job <--- ", j)
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		fmt.Println(" Result <--- ", <-results)
	}

}

// Recebe um id e dois canais que vão controlar a quantidade de jobs e o outro os resultados
func worker(id int, jobs <-chan int, result chan<- int) {
	fmt.Println("w <--- ", id)
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j * 2
	}
}
