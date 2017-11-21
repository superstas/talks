package main

import (
	"fmt"

	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.Create("/tmp/5")
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

	n, err := io.Copy(f, c.Body) // HL
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes was written", n)
	// END 1 OMIT
}
