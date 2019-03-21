package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var sharedVar int64
var sharedVarAtomic int64
var sharedVarMutex int64
var wg sync.WaitGroup
var mutex = &sync.Mutex{}

func addMany(n int, c chan int64) {
	defer wg.Done()
	sharedVar++
	atomic.AddInt64(&sharedVarAtomic, 1)
	mutex.Lock()
	sharedVarMutex++
	mutex.Unlock()
	c <- 1
}

func main() {
	n := 100000
	fmt.Println("Somando de 1 a", n, "usando goroutines...")
	c := make(chan int64, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go addMany(i, c)
	}
	var add int64
	for i := 0; i < n; i++ {
		add += <-c
	}
	wg.Wait()
	fmt.Println("Soma com variável compartilhada:", sharedVar)
	fmt.Println("Soma atômica com variável compartilhada:", sharedVarAtomic)
	fmt.Println("Soma usando mutex em variável compartilhada:", sharedVarMutex)
	fmt.Println("Soma usando channels:", add)
}
