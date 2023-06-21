package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func sayHi(name int) {
	fmt.Println("Hi", name)
}

func TestGoroutine(t *testing.T) {
	runtime.GOMAXPROCS(2) // menentukan banyaknya thread yang akan dipakai
	//go sayHi(230)
	//fmt.Println("Hallo Ali")

	for i := 0; i < 3; i++ {
		go sayHi(i)
	}

	time.Sleep(1 * time.Second)
}

func TestGOMAXPROCS(t *testing.T) {
	thrd := runtime.GOMAXPROCS(20)
	var wg sync.WaitGroup
	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(2 * time.Second)
			wg.Done()
		}()
	}

	fmt.Println("thread:", thrd)

	cpu := runtime.NumCPU()
	fmt.Println("cpu:", cpu)

	goroutine := runtime.NumGoroutine()
	fmt.Println("goroutine:", goroutine)

	wg.Wait()
}
