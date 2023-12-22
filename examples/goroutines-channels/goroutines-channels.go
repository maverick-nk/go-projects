package main

import (
	"fmt"
	"time"
	)

func main(){
	// Creating a channel
	messages := make(chan string)

	/**
	By default sends and receives 
		block until both the sender and receiver are ready. 
	*/

	// Go routines are thread in execution 
	go func(){messages <- "some message"}()

	go func(){
		message := <- messages
		fmt.Println(message)
	}()


	// Asycn go calls so wait for all routines to finish
	time.Sleep(time.Second)
    fmt.Println("Ran all the go routines")


	// By default channels are unbuffered
	// Below is a channel with 2 message buffer
	channel := make(chan string, 2)
	channel <- "first msg"
	channel <- "second msg"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}