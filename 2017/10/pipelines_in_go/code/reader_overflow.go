package main

import (
	"io/ioutil"
)

// ReaderOverflow returns more than len(b) bytes
type ReaderOverflow struct{}

// 0 OMIT
func (r *ReaderOverflow) Read(p []byte) (int, error) {
	//return len(p) + 1, nil
	return -1, nil // HL
}
// END 0 OMIT

func main() {
	r := new(ReaderOverflow)
	ioutil.ReadAll(r)
}
