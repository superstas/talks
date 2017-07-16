package main

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/superstas/future_talks/2017/pipelines_in_go/code/pipelines/readers"
)

func main() {
	// 1 OMIT
	simpleReader := readers.SimpleReader("Hello GoWayFest")
	brokenReader := readers.BrokenReader(simpleReader)
	limitReader := io.LimitReader(brokenReader, 5) // HL
	res, err := ioutil.ReadAll(limitReader)
	fmt.Printf("Result: %s\nError: %v\n", res, err)
	// END 1 OMIT
}
