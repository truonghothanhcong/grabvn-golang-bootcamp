package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

	fmt.Println(result)

	return result, nil
}

func main() {
	countFrequencyAppears("a.txt")
}
