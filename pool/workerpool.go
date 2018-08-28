package main

import (
	"fmt"
	"runtime"
	"sync"
)

func worker(id int, wg *sync.WaitGroup, jobs <-chan int) {
	wg.Add(1)
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		fmt.Println("worker", id, "finished job", j)
		runtime.Gosched()
	}
}

func main() {

	jobs := make(chan int, 3)

	wg := new(sync.WaitGroup)
	for w := 1; w <= 3; w++ {
		go worker(w, wg, jobs)
	}

	for j := 1; j <= 3000; j++ {
		jobs <- j
	}

	close(jobs)

	wg.Wait()

}
