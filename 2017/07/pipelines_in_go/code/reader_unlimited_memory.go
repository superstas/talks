package main

import (
	"io/ioutil"
)

// ReaderUnlimitedMemory always returns 1, nil;
type ReaderUnlimitedMemory struct{}

func (r *ReaderUnlimitedMemory) Read(p []byte) (int, error) {
	return 1, nil // HL
}

func main() {
	r := new(ReaderUnlimitedMemory)
	ioutil.ReadAll(r)
}
