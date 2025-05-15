package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files!")
	content := "This needs to go in a file - LearningGo"

	file, err := os.Create("./mygofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("Length of the file is:", length)
	defer file.Close()
	readFile("./mygofile.txt")
}

func readFile(fname string) {

	databyte, err := ioutil.ReadFile(fname)
	checkNilErr(err)
	fmt.Println("Content in the File:\n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
