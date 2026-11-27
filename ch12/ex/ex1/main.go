package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	var producers sync.WaitGroup
	var consumer sync.WaitGroup

	producers.Add(2)
	consumer.Add(1)

	go func() {
		defer producers.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer producers.Done()
		for i := 20; i < 30; i++ {
			ch <- i
		}
	}()

	go func() {
		defer consumer.Done()
		for value := range ch {
			fmt.Println(value)
		}
	}()

	producers.Wait()
	// we have to close the channel before we can expect the consumer to exit.
	// Because the for-range loop terminates when the channel is CLOSED AND DRAINED.
	close(ch)
	consumer.Wait()
}
