package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"os"
	"path/filepath"
)

var global_result = make(map[string]int)
var numberOfJobs int = 3
var numberOfBufferSizes int = 16

func readAllFilePaths(pathFolder string) ([]string, error) {
	var filePaths []string
	err := filepath.Walk(pathFolder,
		func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filePaths = append(filePaths, path)
		}
		return nil
	})

	return filePaths, err
}

func countFrequencyAppears(filePath string) (map[string]int, error) {
	var result = make(map[string]int)

	fileByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		return result, err
	}
	
	text := string(fileByte)
	wordList := strings.Fields(text)
	for _, w := range wordList {
		if _, ok := result[w]; ok {
			result[w]++
		} else {
			result[w] = 1
		}
	}

	return result, nil
}

func worker(id int, wg *sync.WaitGroup, job <-chan string, writeResult chan<- map[string]int) {
	for path := range job {
		result_count, err := countFrequencyAppears(path)
		if err != nil {
			fmt.Println("error: ", err)
		} else {
			writeResult <- result_count
		}
	}
	wg.Done()
}

func writer(writeResult <-chan map[string]int) {
	for result := range writeResult {
		for key, value := range result {
			if _, ok := global_result[key]; ok {
				global_result[key] += value
			} else {
				global_result[key] = value
			}
		}
	}
}

// TODO: serve for large file by read file in block
// not run now :)
func countFrequencyAppears2(filePath string) (map[string]int, error) {
	var result = make(map[string]int)
	var buffer = make([]byte, numberOfBufferSizes)

	f, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer f.Close()
	
	for {
		_, err := f.Read(buffer)
		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				break
			}
		} else {
			text := string(buffer)
			fmt.Println(text)
		}
	}
	
	// text := string(fileByte)
	// wordList := strings.Fields(text)
	// for _, w := range wordList {
	// 	if _, ok := result[w]; ok {
	// 		result[w]++
	// 	} else {
	// 		result[w] = 1
	// 	}
	// }

	return result, nil
}

func main() {
	jobs := make(chan string)
	result := make(chan map[string]int)
	var wg sync.WaitGroup

	paths, err := readAllFilePaths("./file_folder")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	go writer(result)
	for w := 1; w <= numberOfJobs; w++ {
		wg.Add(1)
		go worker(w, &wg, jobs, result)
	}

	for _, j := range paths {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	fmt.Println(global_result)
}
