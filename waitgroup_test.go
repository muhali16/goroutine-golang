package main

import (
	"encoding/json"
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

	for i := 0; i < 2; i++ {
		data1 := fmt.Sprintf("data %d", i)
		data2 := fmt.Sprintf("data tambahan %d", i+1)

		wg.Add(2)
		go printText(&wg, data1)
		go printText(&wg, data2)

	}

	wg.Wait()
	fmt.Println("End...")
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TestSimpleWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	channel := make(chan User, 1)
	var data User
	jsonString := `{"username": "muhali16", "password": "123", "email": "muhali16@ali.my.id"}`

	jsonToMap := func(data string, ch chan<- User, wg *sync.WaitGroup) {
		defer wg.Done()
		jsonString := []byte(data)
		jsonData := User{}
		err := json.Unmarshal(jsonString, &jsonData)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		ch <- jsonData
	}

	wg.Add(1)
	go jsonToMap(jsonString, channel, &wg)

	data = <-channel
	if data.Password == "" {
		fmt.Println("Password cannot be null!")
		return
	}

	//time.Sleep(3 * time.Second)
	fmt.Println("Username:", data.Username)
	fmt.Println("Email:", data.Email)
	fmt.Println("Password:", data.Password)
	wg.Wait()

}
