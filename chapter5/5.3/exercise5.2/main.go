package main

import (
	"fmt"
	"sync"
	"time"
)

func playerHandler(cond *sync.Cond, playerInGame *int, playerID int, cancel *bool) {
	cond.L.Lock()
	fmt.Println(playerID, ": Connected")
	*playerInGame--
	if *playerInGame == 0 {
		cond.Broadcast()
	}
	for *playerInGame > 0 && !*cancel {
		fmt.Println(playerID, ": Waiting for more players")
		cond.Wait()
	}
	cond.L.Unlock()
	if *cancel {
		fmt.Println(playerID, ": Game cancelled")
	} else {
		fmt.Println("All players connected. Ready player ", playerID)
	}
}

func validateTime(cond *sync.Cond, cancel *bool, t int64) {
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("Quá giờ")
	cond.L.Lock()
	*cancel = true
	cond.Broadcast()
	cond.L.Unlock()
}

func main() {
	playerInGame := 4
	cancel := false
	cond := sync.NewCond(&sync.Mutex{})
	go validateTime(cond, &cancel, 3)
	for i := 0; i < 4; i++ {
		go playerHandler(cond, &playerInGame, i, &cancel)
		time.Sleep(1 * time.Second)
	}
	time.Sleep(60 * time.Second)
}
