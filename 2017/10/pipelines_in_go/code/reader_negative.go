package main

import (
	"io/ioutil"
)

// NegativeReader returns more than len(b) bytes
type NegativeReader struct{}

// 0 OMIT
func (r *NegativeReader) Read(p []byte) (int, error) {
	return -1, nil // HL
}

// END 0 OMIT

func main() {
	// 1 OMIT
	ioutil.ReadAll(new(NegativeReader)) // HL
	// END 1 OMIT
}
