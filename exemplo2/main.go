package main

import (
	"fmt"
	"strconv"
	"time"
)

func dizer(num int, palavra string, canal chan string) {
	time.Sleep(time.Millisecond * 1100)
	canal <- palavra + strconv.Itoa(num)
}

func main() {
	max := 3
	canal1 := make(chan string, max)
	canal2 := make(chan string, max)
	for i := 1; i <= max; i++ {
		go dizer(i, "Olá do canal 1! - Msg ", canal1)
		go dizer(i, "Hello from channel 2! - Msg ", canal2)
	}

	for i := 1; i <= max*3; i++ {
		select {
		case msg := <-canal1:
			fmt.Println(i, "-", msg)
		case msg := <-canal2:
			fmt.Println(i, "-", msg)
		default:
			fmt.Println(i, "- No message!")
			time.Sleep(time.Second * 1)
		}
	}
}
