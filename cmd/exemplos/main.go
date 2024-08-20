package main

import (
	"fmt"

	"time"
)

func count(n int) {
	for i := range n {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {

	ch := make(chan string)

	go func() {
		ch <- "Teste"
	}()

	msg := <-ch
	fmt.Println(msg)
	
}


