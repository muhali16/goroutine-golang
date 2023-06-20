package main

import (
	"fmt"
	"runtime"
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
