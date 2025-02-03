package main

import (
	"fmt"
	"sync"
	"time"
)

type WaitGrp struct {
	groupSize int
	cond      *sync.Cond
}

func NewWaitGrp() *WaitGrp {
	return &WaitGrp{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (wg *WaitGrp) Add(delta int) {
	wg.cond.L.Lock()
	wg.groupSize += delta
	wg.cond.L.Unlock()
}

func (wg *WaitGrp) Wait() {
	wg.cond.L.Lock()
	for wg.groupSize > 0 {
		wg.cond.Wait()
	}
	wg.cond.L.Unlock()
}

func (wg *WaitGrp) Done() {
	wg.cond.L.Lock()
	wg.groupSize--
	if wg.groupSize == 0 {
		wg.cond.Broadcast()
	}
	wg.cond.L.Unlock()
}

func DoWork(i int, wg *sync.WaitGroup) {
	fmt.Println("Doing task: ", i)
	time.Sleep(1 * time.Second)
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(15)
	for i := 0; i < 15; i++ {
		go DoWork(i, &wg)
	}
	wg.Wait()
	fmt.Println("All job is done")
}
