package main

import (
	"fmt"
	"sync"

)

func main() {
	var mu sync.Mutex
	a := 0 // data race
	//done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			a += 1
			wg.Done()
			//done <- struct{}{}
			
		}()
	}
	wg.Wait()
	//<- done
  	fmt.Println(a) // could theoretical be any number 0-1000 (most likely above 900)
}