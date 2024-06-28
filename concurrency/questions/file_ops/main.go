package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func readFile(filename string, dataChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dataChan <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner unable to read file:", err)
	}
}

func main() {
	var wg sync.WaitGroup
	dataChan := make(chan string, 2)

	directoryName := "files"

	files, err := os.ReadDir(directoryName)
	if err != nil {
		fmt.Println("Error reading from directory:", err)
	}

	for _, file := range files {
		wg.Add(1)
		go readFile(filepath.Join(directoryName, file.Name()), dataChan, &wg)
	}

	commonFile, err := os.Create("commonFile.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer commonFile.Close()

	writer := bufio.NewWriter(commonFile)
	for line := range dataChan {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writting to common file:", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	commonData, err := os.ReadFile("commonFile.txt")
	if err != nil {
		fmt.Println("Unable to read from common file:", err)
		return
	}
	fmt.Println("Content of Common file: ", string(commonData))

	go func() {
		wg.Wait()
		close(dataChan)
	}()

}
