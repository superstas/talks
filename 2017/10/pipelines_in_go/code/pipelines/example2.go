package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.Create("/tmp/3")
	defer f.Close()
	if err != nil {
		log.Fatal("failed to create file: ", err)
	}

	// 1 OMIT
	c, err := http.Get("https://golang.org/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Body.Close()

	r := io.TeeReader(c.Body, f) // HL
	resp, err := ioutil.ReadAll(r)
	fmt.Printf("###\n%s###\n", resp)
	// END 1 OMIT
}
