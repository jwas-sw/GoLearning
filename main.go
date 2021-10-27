package main

import (
	"fmt"
	"os"
	"github.com/eiannone/keyboard"
	eh "github.com/jwas-sw/GoLearning/v2/errorHandler"
	fh "github.com/jwas-sw/GoLearning/v2/fileHandler"
	fp "github.com/jwas-sw/GoLearning/v2/fileProcessor"
)

func main() {
	for {
		buildMenu()
	}

	// TODO
	// unit tests
}

func buildMenu() {
	fmt.Println("Select mode: Results merge mode[2] or Text analysis mode[1]. Press 'X' to exit ")
	char, _, err := keyboard.GetSingleKey()
	eh.Check(err)
	fmt.Printf("You pressed: %q\r\n", char)
	switch char {
	case '1':
		fmt.Println("Text analysis mode seleted.")
		byteArrayChannel1 := make(chan []byte)
		byteArrayChannel2 := make(chan []byte)

		go fp.DownloadFromUrl("https://pastebin.com/raw/v0Sm2sfn", byteArrayChannel1)
		go fp.DownloadFromUrl("https://pastebin.com/raw/fysHJ7YX", byteArrayChannel2)

		fh.SaveToFile("output1.txt", <-byteArrayChannel1)
		fh.SaveToFile("output2.txt", <-byteArrayChannel2)

        fp.CreateJsonFromFile("output1.txt", "output1.json")
		fp.CreateJsonFromFile("output2.txt", "output2.json")
	case '2':
		fmt.Println("Results merge mode mode selected.")
		mergedFilesMap := fp.MergeFilesIntoMap("output1.json", "output2.json")
		jsonByteMap := fp.CreateJson(mergedFilesMap)
		fh.SaveToFile("output3.json", jsonByteMap)
    case '3':
        fmt.Println("Printing outputs mode...")
        fp.PrintFileStatistics("output1.json")
        fp.PrintFileStatistics("output2.json")
        fp.PrintFileStatistics("output3.json")
	default:
		fmt.Println("Unsuported key pressed. Killing...")
		os.Exit(0)
	}
}
