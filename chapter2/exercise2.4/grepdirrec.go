package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"
)

// Adapt the program in the third exercise to continue searching recursively
// in any subdirectories. If you give your search goroutine a file, it should
// search for a string match in that file, just like in the previous exercises.
// Otherwise, if you give it a directory, it should recursively spawn a new goroutine
// for each file or directory found inside. Call the program grepdirrec.go, and
// execute it by running this command:
// go run chapter2/exercise2.4/grepdirrec.go abc chapter2/

func grepDirRecursive(fileDir string, fileName fs.DirEntry, matchWord string) {
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
	} else {
		files, err := os.ReadDir(filePath)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			go grepDirRecursive(filePath+"/", file, matchWord)
		}
	}
}

func main() {
	matchWord := os.Args[1]
	fileDir := os.Args[2]
	files, err := os.ReadDir(fileDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileName := range files {
		go grepDirRecursive(fileDir, fileName, matchWord)
	}
	time.Sleep(2 * time.Second)
}
