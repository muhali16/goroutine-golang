package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)

	sayName := func(name string) {
		data := fmt.Sprintf("%s", name)
		channel <- data
	}

	go sayName("Ali")
	go sayName("Loren")

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	close(channel)
}

func TestChannelAsParameter(t *testing.T) {
	outCar := func(car <-chan string) { // receive channel data function
		fmt.Println("name:", <-car)
		fmt.Println("tire:", <-car)
		fmt.Println("type:", <-car)
	}
	inCar := func(car chan<- string) { // send data to channel param function
		for _, cars := range []string{"Toyota", "4", "Car"} {
			car <- cars
		}
	}

	channel := make(chan string)
	//defer close(channel)

	go inCar(channel)
	outCar(channel)
}

func TestBufferedChannel(t *testing.T) {
	buffer := make(chan int, 2)

	go func() {
		for {
			i := <-buffer
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		buffer <- i
		fmt.Println("Send data", i)
	}
	time.Sleep(2 * time.Second)
}

func TestSelectChannel(t *testing.T) {
	runtime.GOMAXPROCS(2)
	getMax := func(num []int, ch chan int) {
		var data = 0
		for _, each := range num {
			if each > data {
				data = each
			}
		}

		ch <- data
	}
	getAverage := func(num []int, ch chan float64) {
		var data float64
		for i := 0; i < len(num); i++ {
			data += float64(num[i])
		}
		ch <- data / float64(len(num))
	}

	number := []int{12, 23, 45, 23, 12, 64, 24, 76, 22, 56}

	ch1 := make(chan int)
	go getMax(number, ch1)

	ch2 := make(chan float64)
	go getAverage(number, ch2)

	select { // print wether chann that send data
	case max := <-ch1:
		fmt.Println("Max:", max)
	case avg := <-ch2:
		fmt.Println("Average:", avg)
	}
}

func TestTimeoutChannel(t *testing.T) {
	sendMessage := func(ch chan<- int) {
		for i := 0; true; i++ {
			ch <- i
			time.Sleep(time.Duration(rand.Int()%10) * time.Second)
		}
	}
	receiveMessage := func(ch <-chan int) {
	ulang:
		for {
			select {
			case data := <-ch:
				fmt.Println("receive data", data)
			case <-time.After(5 * time.Second):
				fmt.Println("Timeout 5 second ...")
				break ulang
			}
		}
	}

	message := make(chan int)
	go sendMessage(message)
	receiveMessage(message)
}
