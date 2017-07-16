package main

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/superstas/future_talks/2017/pipelines_in_go/code/pipelines/readers"
)

func main() {
	// 1 OMIT
	r := readers.BrokenReader(io.LimitReader(readers.SimpleReader("Hello GoWayFest"), 5)) // HL
	res, err := ioutil.ReadAll(r)
	fmt.Printf("Result: %s\nError: %v\n", res, err)
	// END 1 OMIT
}
