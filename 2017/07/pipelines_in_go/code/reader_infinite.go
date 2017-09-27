package main

import (
	"io/ioutil"
)

// ReaderInfinite always returns 0, nil;
type ReaderInfinite struct{}

func (r *ReaderInfinite) Read(p []byte) (int, error) {
	return 0, nil // HL
}

func main() {
	r := new(ReaderInfinite)
	ioutil.ReadAll(r)
}
