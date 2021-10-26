package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"github.com/eiannone/keyboard"
)

type tmpMap struct {
	i int
	s string
}

func main() {
	buildMenu()
	//todo

	//select top 50 most used words and save them into jsons (!!!)

	// modules? packages?
	// unit tests
	// error handling?
}

func buildMenu() {
	fmt.Println("Select mode: Results merge mode[0] or Text analysis mode[1]. Press 'X' to exit ")
	char, _, err := keyboard.GetSingleKey()
	check(err)
	fmt.Printf("You pressed: %q\r\n", char)
	switch char {
	case '1':
		fmt.Println("Text analysis mode seleted.")
        byteArrayChannel1 := make(chan []byte)
        byteArrayChannel2 := make(chan []byte)

        go downloadFromUrl("https://pastebin.com/raw/v0Sm2sfn", byteArrayChannel1)
		go downloadFromUrl("https://pastebin.com/raw/fysHJ7YX", byteArrayChannel2)

		saveToFile("output1.txt", <-byteArrayChannel1)
		saveToFile("output2.txt", <-byteArrayChannel2)
	case '0':
		fmt.Println("Results merge mode mode selected.")

		createJsonFromFile("output1.txt", "output1.json")
		createJsonFromFile("output2.txt", "output2.json")

		mergedFilesMap := mergeFilesIntoMap("output1.txt", "output2.txt")
		jsonByteMap := createJsonByteFromMap(mergedFilesMap)
		saveToFile("output3.json", jsonByteMap)

	default:
		fmt.Println("Unsuported key pressed")
		os.Exit(0)
	}
}

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

func saveToFile(outputFileName string, contentAsByte []byte) {
	f, err := os.Create(outputFileName)
	// close fo on exit and check for its returned error
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	check(err)
	_, err2 := f.Write(contentAsByte)
	check(err2)
}

func openFile(fileName string, c chan string) {
	fi, err := ioutil.ReadFile(fileName)
	check(err)
	c <- string(fi)
}

func wordCount(c chan string) (x map[string]int) {
	x = make(map[string]int)
	for _, j := range strings.Fields(<-c) {
		x[j]++
	}

	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
