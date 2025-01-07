package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Write a program similar to the one in listing 2.3 that accepts a list of text file-
// names as arguments. For each filename, the program should spawn a new goroutine that
// will output the contents of that file to the console. You can use the time.Sleep()
// function to wait for the child goroutines to complete (until you know how to do this better).
// Call the program catfiles.go. Here’s how you can execute this Go program:
// go run catfiles.go txtfile1 txtfile2 txtfile3

// func printFileContent(fileName string){
// 	fileContent, err := os.ReadFile(fileName)
// 	if err  != nil

// }

func main() {
	fileName := os.Args[1:]
	for _, file := range fileName {
		go func(fileName string) {
			fileContent, err := os.ReadFile(fileName)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(fileContent))
		}(file)
	}
	time.Sleep(3 * time.Second)
}
