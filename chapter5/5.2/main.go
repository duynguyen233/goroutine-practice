package main

import (
	"fmt"
	"sync"
	"time"
)

func doWork(cond *sync.Cond) {
	fmt.Println("Work Start")
	fmt.Println("Work Ended")
	cond.L.Lock()
	fmt.Println("Lock release")
	cond.Signal()
	cond.L.Unlock()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	cond.L.Lock()
	for i := 0; i < 1; i++ {
		go doWork(cond)
		fmt.Println("Waiting for child goroutine ")
		time.Sleep(1 * time.Second)
		cond.Wait()
		fmt.Println("Child goroutine finished")
	}
	cond.L.Unlock()
}
