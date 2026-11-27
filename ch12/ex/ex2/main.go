package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var consumerWg sync.WaitGroup

	consumerWg.Add(1)

	go func() {
		defer close(ch1)
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 30; i < 40; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer consumerWg.Done()
		for finishedCnt := 0; finishedCnt < 2; {
			select {
			case current, ok := <-ch1:
				if !ok {
					ch1 = nil
					finishedCnt++
					// we the channel is closed, we shouldn't print a value
					// which is why we use continue here
					continue
				}
				fmt.Println("Goroutine 1:", current)

			case current, ok := <-ch2:
				if !ok {
					ch2 = nil
					finishedCnt++
					// we the channel is closed, we shouldn't print a value
					// which is why we use continue here
					continue
				}
				fmt.Println("Gorountine 2:", current)
			}
		}
	}()

	consumerWg.Wait()
}
