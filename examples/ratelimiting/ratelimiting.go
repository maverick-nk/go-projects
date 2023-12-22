package main

import (
	"fmt"
	"time"
)

func main(){

	requests := make(chan int, 5)

	for i:=0;i<5;i++ {
	
		requests <- i // sending requests
	}
	close(requests)

	ticker := time.Tick(200 * time.Millisecond)	

	for req := range requests {
		<-ticker
		fmt.Println("Requests", req, time.Now())
	}

}