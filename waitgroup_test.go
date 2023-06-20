package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	runtime.GOMAXPROCS(2)
	printText := func(wg *sync.WaitGroup, message string) {
		defer wg.Done()
		fmt.Println(message)
	}
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		data1 := fmt.Sprintf("data %d", i)
		data2 := fmt.Sprintf("data tambahan %d", i+1)

		wg.Add(2)
		go printText(&wg, data1)
		go printText(&wg, data2)

	}

	wg.Wait()
	fmt.Println("End...")
}
