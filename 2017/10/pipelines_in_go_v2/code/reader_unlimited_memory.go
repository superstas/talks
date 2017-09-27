package main

import (
	"io/ioutil"
)

// ReaderUnlimitedMemory always returns 1, nil;
type ReaderUnlimitedMemory struct{}

// 0 OMIT
func (r *ReaderUnlimitedMemory) Read(p []byte) (int, error) {
	return 1, nil // HL
}
// END 0 OMIT


func main() {
	r := new(ReaderUnlimitedMemory)
	ioutil.ReadAll(r)
}
