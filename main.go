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
}

func buildMenu() {
    fmt.Println("What I can do:")
    fmt.Println("1 - Text analysis mode")
    fmt.Println("2 - Merge results")
    fmt.Println("3 - Print all result files")
	fmt.Println("4 - Remove all created files")
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
        fmt.Println("Done.")
	case '2':
		fmt.Println("Results merge mode mode selected.")
		mergedFilesMap := fp.MergeFilesIntoMap("output1.json", "output2.json")
		jsonByteMap := fp.CreateJson(mergedFilesMap)
		fh.SaveToFile("output3.json", jsonByteMap)
        fmt.Println("Done.")
    case '3':
        fmt.Println("Printing outputs mode...")
        fp.PrintFileStatistics("output1.json")
        fp.PrintFileStatistics("output2.json")
        fp.PrintFileStatistics("output3.json")
        fmt.Println("Done.")
    case '4':
        fh.DeleteCreatedFiles()
        fmt.Println("Done.")
	default:
		fmt.Println("Unsuported key pressed. Killing...")
		os.Exit(0)
	}
}
