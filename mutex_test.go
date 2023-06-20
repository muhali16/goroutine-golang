package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

type Counter struct {
	Val int
}

func (c *Counter) Count(int) {
	c.Val++
}

func (c *Counter) Value() int {
	return c.Val
}

func TestRaceCondition(t *testing.T) {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	var count Counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				count.Count(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count.Value())
}

func TestMutex(t *testing.T) {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	var mtx sync.Mutex
	var count Counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				mtx.Lock()
				count.Count(1)
				mtx.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count.Value())
}
