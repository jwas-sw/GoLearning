package fileProcessor

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
    eh "github.com/jwas-sw/GoLearning/v2/errorHandler"
    fh "github.com/jwas-sw/GoLearning/v2/fileHandler"
)

func CreateJsonByteFromMap(someMap map[string]int) []byte {
	jsonStr, err := json.Marshal(someMap)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}
	return jsonStr
}

func CreateMapFromFile(fileName string) (x map[string]int) {
	c := make(chan string)
	go fh.OpenFile(fileName, c)
	m := WordCount(c)
	return m
}

func CreateJsonFromFile(inputFileName, outputFileName string) {
	m := CreateMapFromFile(inputFileName)
	jsonByteMap := CreateJsonByteFromMap(m)
	fh.SaveToFile(outputFileName, jsonByteMap)
}

func MergeFilesIntoMap(filename1, fileName2 string) (x map[string]int) {
	c := make(chan string)
	go fh.OpenFile(filename1, c)
	go fh.OpenFile(fileName2, c)
	m := WordCount(c)
	return m
}

func DownloadFromUrl(url string, c chan []byte) {
	resp, err := http.Get(url)
	eh.Check(err)
	// close fo on exit and check for its returned error
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()
	result, err := io.ReadAll(resp.Body)
	eh.Check(err)
	c <- result
}

func WordCount(c chan string) (x map[string]int) {
	x = make(map[string]int)
	for _, j := range strings.Fields(<-c) {
		x[j]++
	}

	return
}
