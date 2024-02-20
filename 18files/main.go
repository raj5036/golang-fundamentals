package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files handling in Golang")
	var content string = "This needs to go in a file"
	file, err := os.Create("./myFile")

	if err != nil {
		panic(err)
	} else {
		length, errInWriting := io.WriteString(file, content)
		if errInWriting != nil {
			panic(errInWriting)
		} else {
			fmt.Println("Length is ", length)
			defer file.Close() // recommended way is to use 'defer'
		}
	}
	readFile("./myFile")
}

func readFile(filename string) {
	dataBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Data inside file is ", dataBytes)
		fmt.Println("Data is ", string(dataBytes))
	}
}
