package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x:= range data {
		fmt.Printf("worker %d got %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {

	ch := make(chan int)
	workers := 9
	for i := range workers {
		
		go worker(i, ch)
	}
	
	
	for i := range 20 {
		ch <- i
	}
}
