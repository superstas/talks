package main

import (
	"io/ioutil"
)

// ReaderOverflow returns more than len(b) bytes
type ReaderOverflow struct{}

func (r *ReaderOverflow) Read(p []byte) (int, error) {
	return len(p) + 1, nil // HL
	//return -1, nil
}

func main() {
	r := new(ReaderOverflow)
	ioutil.ReadAll(r)
}
