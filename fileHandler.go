package fileHandler

import (
    "github.com/jwas-sw/GoLearning/v2/errorHandler"
)

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
