/*A buffered channel in Go is a channel that has a buffer, which is a queue of values that can be stored in the channel.
When you send a value to a buffered channel,
it's stored in the buffer until it's received by another goroutine. */

// ch := make(chan int, 10) // declares a buffered channel of type int with a buffer size of 10

package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Producing %d\n", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch chan int) {
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Printf("Consumed %d\n", msg)
		}
	}
}

func main() {
	ch := make(chan int, 5) // buffered channel with a buffer size of 5
	go producer(ch)
	go consumer(ch)
	time.Sleep(2 * time.Second)
}
