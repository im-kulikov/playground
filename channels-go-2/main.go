package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var tmp = make([]<-chan int, 10)
	for i := 0; i < 10; i++ {
		var ints []int
		for j := i * 10; j < i*10+10; j++ {
			ints = append(ints, j)
		}
		tmp[i] = asChan(ints...)
	}
	g := merge(tmp...)
	for v := range g {
		fmt.Println(v)
	}
}

func merge(channels ...<-chan int) chan int {
	var wg sync.WaitGroup
	var result = make(chan int)
	for i := range channels {
		wg.Add(1)
		go func(input <-chan int, num int) {
			for val := range input {
				result <- val
			}
			fmt.Printf("Channel done: %d\n", num)
			wg.Done()
		}(channels[i], i)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}
