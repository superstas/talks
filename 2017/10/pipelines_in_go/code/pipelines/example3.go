package main

import (
	"fmt"

	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.Create("/tmp/4")
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

	// wrong way :(
	buf := make([]byte, 128)
	n, err := c.Body.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	// there is a better way
	fmt.Printf("%d bytes was read\n", n)

	n, err = f.Write(buf[:n])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes was written", n)
	// END 1 OMIT
}
