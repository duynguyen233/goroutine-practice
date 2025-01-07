package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"
)

// Change the program you wrote in the second exercise so that instead of passing
// a list of text filenames, you pass a directory path. The program will look inside
// this directory and list the files. For each file, you can spawn a goroutine that
// will search for a string match (the same as before). Call the program grepdir.go.
// Here’s how you can execute this Go program:
// go run chapter2/exercise2.3/grepdir.go abc chapter2/

func main() {
	matchWord := os.Args[1]
	fileDir := os.Args[2]
	files, err := os.ReadDir(fileDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileName := range files {
		go func(fileName fs.DirEntry) {
			filePath := fileDir + fileName.Name()
			if !fileName.IsDir() {
				fileContent, err := os.ReadFile(filePath)
				if err != nil {
					log.Fatal(err)
				}
				if strings.Contains(string(fileContent), matchWord) {
					fmt.Println(matchWord, "contains in", fileName.Name())
				} else {
					fmt.Println(matchWord, "does not contain in", fileName.Name())
				}
			}
		}(fileName)
	}
	time.Sleep(2 * time.Second)
}
