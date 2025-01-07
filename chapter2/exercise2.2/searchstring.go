package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Expand the program you wrote in the first exercise so that instead of printing
// the contents of the text files, it searches for a string match. The string to
// search for is the first argument on the command line. When you spawn a new goroutine,
// instead of printing the file’s contents, it should read the file and search for a match.
// If the goroutine finds a match, it should output a message saying that the filename contains a match.
// Call the program grepfiles.go. Here’s how you can execute this Go program (“bubbles” is the search
// string in this example):
// go run chapter2/exercise2.2/searchstring.go abc chapter2/txtfile1 chapter2/txtfile2 chapter2/txtfile3

func main() {
	matchWord := os.Args[1]
	file := os.Args[2:]
	for _, fileName := range file {
		go func(fileName string) {
			content, err := os.ReadFile(fileName)
			if err != nil {
				log.Fatal("Failed to read file ", err)
			}
			if strings.Contains(string(content), matchWord) {
				fmt.Println(matchWord, "contains in file", fileName)
			} else {
				fmt.Println(fileName, "not contains the word", matchWord)
			}
		}(fileName)
	}
	time.Sleep(2 * time.Second)
}
