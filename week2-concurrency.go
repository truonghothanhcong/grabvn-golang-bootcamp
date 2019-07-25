package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

func readPathFromFolder(pathFolder string) ([]string, error) {
	var filePaths []string
	filesInfo, err := ioutil.ReadDir(pathFolder)
	if err != nil {
		return filePaths, err
	}

	for _, f := range filesInfo {
		filePaths = append(filePaths, pathFolder + f.Name())
	}

	return filePaths, nil
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

var global_result = make(map[string]int)

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

func main() {
	jobs := make(chan string)
	result := make(chan map[string]int)
	var wg sync.WaitGroup

	paths, err := readPathFromFolder("./file_folder/")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(paths)
	// for _, j := range paths {
	// 	result_1, _ := countFrequencyAppears(j)
	// 	fmt.Println(result_1)
	// }
	
	go writer(result)
	for w := 1; w <= 3; w++ {
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
