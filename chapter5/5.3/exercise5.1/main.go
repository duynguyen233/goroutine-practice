package main

import (
	"fmt"
	"os"
	"sync"
)

// In listing 5.4, Stingy’s goroutine is signaling on the condition variable every
// time we add money to the bank account. Can you change the function so that it signals
// only when there is $50 or more in the account?
func stingy(money *int, cond *sync.Cond) {
	for i := 0; i < 1000000; i++ {
		cond.L.Lock()
		*money += 10
		if *money > 50 {
			cond.Signal()
		}
		cond.L.Unlock()
	}
	fmt.Println("Stingy done")
}

func spendy(money *int, cond *sync.Cond) {
	for i := 0; i < 200000; i++ {
		cond.L.Lock()
		for *money < 50 {
			cond.Wait()
		}
		*money -= 50
		if *money < 0 {
			fmt.Println("Money is negative")
			os.Exit(1)
		}
		cond.L.Unlock()
	}
	fmt.Println("Spendy done")
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	money := 100
	done := sync.WaitGroup{}
	done.Add(2)
	go func() {
		defer done.Done()
		stingy(&money, cond)
	}()
	go func() {
		defer done.Done()
		spendy(&money, cond)
	}()
	done.Wait()
	fmt.Println("Money is: ", money)
}
