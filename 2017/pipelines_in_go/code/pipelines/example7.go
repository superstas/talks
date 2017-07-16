package main

import (
	"github.com/superstas/future_talks/2017/pipelines_in_go/code/pipelines/readers"
	"github.com/superstas/future_talks/2017/pipelines_in_go/code/pipelines/writers"
	"io"
	"fmt"
	"io/ioutil"
	"bufio"
)

func main() {
	// 1 OMIT
	r := readers.SimpleReader("Hello GoWayFest")
	w := writers.SimpleWriter("GoWayFestWriter")
	// (reader1 | writer1) - teeReader
	tee := io.TeeReader(r, w)

	// teeReader|limitReader
	lr := io.LimitReader(bufio.NewReader(tee), 10) // HL

	res, err := ioutil.ReadAll(lr)
	fmt.Printf("Result (%d bytes): %s\nError: %v\n", len(res), res, err)
	// END 1 OMIT
}
