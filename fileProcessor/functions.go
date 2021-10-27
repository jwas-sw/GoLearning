package fileProcessor

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	eh "github.com/jwas-sw/GoLearning/v2/errorHandler"
	fh "github.com/jwas-sw/GoLearning/v2/fileHandler"
)

func PrintFileStatistics(fileName string) {
	fmt.Println("FileName - ", fileName)
	fmt.Println(CreateFileStatistics(fileName))
	fmt.Println()
}

func CreateJson(wordCount []StringInt) []byte {
	jsonStr, err := json.Marshal(wordCount)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		//fmt.Println(string(jsonStr))
	}
	return jsonStr
}

func CreateFileStatistics(fileName string) []StringInt {
	c := make(chan string)
	go fh.OpenFile(fileName, c)
	m := WordCount(c)
	return m
}

func CreateJsonFromFile(inputFileName, outputFileName string) {
	m := CreateFileStatistics(inputFileName)
	jsonByteMap := CreateJson(m)
	fh.SaveToFile(outputFileName, jsonByteMap)
}

func MergeFilesIntoMap(filename1, fileName2 string) []StringInt {
	c := make(chan string)
    c2 := make(chan string)

	go fh.OpenFile(filename1, c)
	go fh.OpenFile(fileName2, c2)

    data1 := []StringInt{}
    json.Unmarshal([]byte(<-c), &data1)

    data2 := []StringInt{}
    json.Unmarshal([]byte(<-c2), &data2)

    results := []StringInt{}
    results = FindAndUpdate(data1, results)
    results = FindAndUpdate(data2, results)

    finalResult := SortStringIntSlice(results)

	return finalResult
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

func WordCount(c chan string) []StringInt {
	someMap := make(map[string]int)
	for _, j := range strings.Fields(<-c) {
		someMap[j]++
	}

	var results []StringInt
	for key, element := range someMap {
		results = append(results, StringInt{Word: key, Count: element})
	}

    return SortStringIntSlice(results)
}

func SortStringIntSlice(results []StringInt) []StringInt{
	sort.Slice(results, func(i, j int) bool {
		return results[i].Count > results[j].Count
	})

	if len(results) > 50 {
		results = results[0:49]
	}

	return results
}
