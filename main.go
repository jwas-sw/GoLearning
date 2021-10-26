package main

import (
	"encoding/json"
	"fmt"
	"github.com/eiannone/keyboard"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
    "github.com/jwas-sw/GoLearning/v2/errorHandler"
    "github.com/jwas-sw/GoLearning/v2/fileHandler"
)

func main() {
    for {
        buildMenu()
    }

	//todo

	//select top 50 most used words and save them into jsons (!!!)

	// modules? packages?
	// unit tests
	// error handling?
}

func buildMenu() {
	fmt.Println("Select mode: Results merge mode[2] or Text analysis mode[1]. Press 'X' to exit ")
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
	case '2':
		fmt.Println("Results merge mode mode selected.")

		createJsonFromFile("output1.txt", "output1.json")
		createJsonFromFile("output2.txt", "output2.json")

		mergedFilesMap := mergeFilesIntoMap("output1.txt", "output2.txt")
		jsonByteMap := createJsonByteFromMap(mergedFilesMap)
		saveToFile("output3.json", jsonByteMap)

	default:
		fmt.Println("Unsuported key pressed. Killing...")
		os.Exit(0)
	}
}
