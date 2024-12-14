package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	//defer the waitgroup Done
	defer wg.Done()
	fmt.Printf("Worker %d started\n", i)
	//soem task is happening here
	fmt.Printf("Worker %d end\n", i)
}

//worker (1)
// worker (2)
// worker (3)

func main() {
	// fmt.Println("Explore goroutine started")

	var wg sync.WaitGroup
	// start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) //increment the waitgroup counter
		go worker(i, &wg)
	}

	//wait for the workers to finish
	wg.Wait()
	fmt.Println("workers task complete")

}
