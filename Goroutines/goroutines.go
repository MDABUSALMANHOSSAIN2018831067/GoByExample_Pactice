package main

import (
	"time"
	"fmt"
)

func main() {
	var numbers []int // nil

	done := make(chan struct{})
	

	// start a goroutine to initialise array
	go func () {
		numbers = make([]int, 2)
		done <- struct{}{}
	}()

	<- done
  
  // do something synchronous
  if numbers == nil {
    time.Sleep(time.Second)
  }
  numbers[0] = 1 // will sometimes panic here
	fmt.Println(numbers[0])
	
}