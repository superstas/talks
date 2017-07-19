package main

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/superstas/future_talks/2017/pipelines_in_go/code/pipelines/readers"
	"github.com/superstas/future_talks/2017/pipelines_in_go/code/pipelines/writers"
)

func main() {
	// 1 OMIT
	r := readers.SimpleReader("Hello GoWayFest")
	w := writers.SimpleWriter("GoWayFestWriter")
	// (reader1 | writer1) - teeReader
	tee := io.TeeReader(r, w)

	// teeReader|BrokenReader
	br := readers.BrokenReader(tee) // HL

	res, err := ioutil.ReadAll(br)
	fmt.Printf("Result (%d bytes): %s\nError: %v\n", len(res), res, err)
	// END 1 OMIT
}
