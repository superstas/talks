package utils

import (
	"io/ioutil"
	"log"
	"os"
)

func MakeTmpFile(fileName string) *os.File {
	err := ioutil.WriteFile(fileName, []byte("Hello World!"), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
