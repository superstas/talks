package main

import (
	"fmt"
	"io/ioutil"
	"github.com/superstas/future_talks/2017/pipelines_in_go/code/pipelines/readers"
)


func main() {
	// 3 OMIT
	r := readers.BrokenReader(readers.SimpleReader("Hello GoWayFest"))
	res, err := ioutil.ReadAll(r)
	fmt.Printf("Result: %s\nError: %v\n", res, err)
	// END 3 OMIT
}
