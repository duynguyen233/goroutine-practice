package main

import (
	"fmt"
	"time"
)

func test(n int) {
	fmt.Println(n)
	for i := range n {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("Test is terminated")
}

func main() {
	go test(10)
	time.Sleep(3 * time.Second)
}
