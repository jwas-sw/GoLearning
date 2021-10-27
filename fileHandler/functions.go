package fileHandler

import (
	"io/ioutil"
	"os"
    eh "github.com/jwas-sw/GoLearning/v2/errorHandler"
)

func SaveToFile(outputFileName string, contentAsByte []byte) {
	f, err := os.Create(outputFileName)
	// close fo on exit and check for its returned error
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	eh.Check(err)
	_, err2 := f.Write(contentAsByte)
	eh.Check(err2)
}

func OpenFile(fileName string, c chan string) {
	fi, err := ioutil.ReadFile(fileName)
	eh.Check(err)
	c <- string(fi)
}
