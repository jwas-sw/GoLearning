package fileProcessor

import (
    "github.com/jwas-sw/GoLearning/v2/errorHandler"
)

func createJsonByteFromMap(someMap map[string]int) []byte {
	jsonStr, err := json.Marshal(someMap)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}
	return jsonStr
}

func createMapFromFile(fileName string) (x map[string]int) {
	c := make(chan string)
	go openFile(fileName, c)
	m := wordCount(c)
	return m
}

func createJsonFromFile(inputFileName, outputFileName string) {
	m := createMapFromFile(inputFileName)
	jsonByteMap := createJsonByteFromMap(m)
	saveToFile(outputFileName, jsonByteMap)
}

func mergeFilesIntoMap(filename1, fileName2 string) (x map[string]int) {
	c := make(chan string)
	go openFile(filename1, c)
	go openFile(fileName2, c)
	m := wordCount(c)
	return m
}

func downloadFromUrl(url string, c chan []byte) {
	resp, err := http.Get(url)
	check(err)
	// close fo on exit and check for its returned error
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()
	result, err := io.ReadAll(resp.Body)
	check(err)
	c <- result
}

func wordCount(c chan string) (x map[string]int) {
	x = make(map[string]int)
	for _, j := range strings.Fields(<-c) {
		x[j]++
	}

	return
}
