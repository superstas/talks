package main

import (
	"github.com/superstas/future_talks/2017/pipelines_in_go/code/pipelines/readers"
	"github.com/superstas/future_talks/2017/pipelines_in_go/code/pipelines/writers"
)

func main() {
	// 1 OMIT
	r := readers.SimpleReader("Hello GoWayFest")
	w := writers.SimpleWriter("GoWayFestWriter")
	r.WriteTo(w) // HL
	// END 1 OMIT
}
