package main

import (
	"io/ioutil"
)

// ReaderInfinite always returns 0, nil;
type ReaderInfinite struct{}

// 0 OMIT
func (r *ReaderInfinite) Read(p []byte) (int, error) {
	return 0, nil // HL
}

// END 0 OMIT

func main() {
	// 1 OMIT
	r := new(ReaderInfinite)
	ioutil.ReadAll(r)
	// END 1 OMIT
}
