package main

import (
	"fmt"
	"time"
)

func main(){
	channel1 := make(chan int)
	channel2 := make(chan string)


	go func(){
		time.Sleep(1 * time.Second)
		channel1 <- 1
	}()

	go func(){
		time.Sleep(4 * time.Second)
		channel1 <- 2
		channel2 <- "Message"
	}()

	for i:=0; i<3;i++ {
		select {
			case msg1 := <-channel1:
				fmt.Println(msg1)
			case msg2 := <-channel2:
				fmt.Println(msg2)
		}
	}

	// Closing channel 
	queue  := make(chan int, 5)
	for i:=0;i<5;i++{
		queue <- i
	}
	close(queue) // closing channel

	for msg := range queue {
		fmt.Println(msg)
	}
}