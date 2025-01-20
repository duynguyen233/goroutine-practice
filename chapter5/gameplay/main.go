package main

import (
	"fmt"
	"sync"
	"time"
)

func playerHandler(cond *sync.Cond, playerInGame *int, playerID int) {
	cond.L.Lock()
	fmt.Println(playerID, ": Connected")
	*playerInGame--
	if *playerInGame == 0 {
		cond.Broadcast()
	}
	for *playerInGame > 0 {
		fmt.Println(playerID, ": Waiting for more players")
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Println("All players connected. Ready player ", playerID)
}

func main() {
	playerInGame := 4
	cond := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 4; i++ {
		go playerHandler(cond, &playerInGame, i)
		time.Sleep(1 * time.Second)
	}
}
