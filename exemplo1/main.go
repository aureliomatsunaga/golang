package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func somar(x *int) {
	defer wg.Done()
	*x++
}

func main() {
	var x, y int
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go somar(&x)
	}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		somar(&y)
	}
	wg.Wait()
	fmt.Println("Soma utilizando concorrência ->", x)
	fmt.Println("Soma sem utilizar concorrência ->", y)
}
