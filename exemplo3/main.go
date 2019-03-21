package main

import (
	"fmt"
	"time"
)

func dizer(num int, palavra string, canal chan string) {
	time.Sleep(time.Second * 4)
	canal <- palavra
}

func main() {
	for i := 1; i <= 6; i++ {
		start := time.Now()
		canal := make(chan string)
		go dizer(i, "- Hello", canal)
		select {
		case msg := <-canal:
			fmt.Println(i, msg)
		case <-time.After(time.Second * time.Duration(i)):
			fmt.Println(i, "- Timeout")
		}
		fmt.Println("Time elapsed", time.Since(start))
	}
}
