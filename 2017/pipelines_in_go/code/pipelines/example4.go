package main

import (
	"github.com/superstas/future_talks/2017/pipelines_in_go/code/pipelines/readers"
	"github.com/superstas/future_talks/2017/pipelines_in_go/code/pipelines/writers"
)

func main() {
	// 1 OMIT
	simpleReader := readers.SimpleReader("Hello GoWayFest")
	simpleWriter := writers.SimpleWriter("GoWayFestWriter")
	buf := make([]byte, 16) // redundant buffer // HL
	simpleReader.Read(buf)
	simpleWriter.Write(buf)
	// END 1 OMIT
}
