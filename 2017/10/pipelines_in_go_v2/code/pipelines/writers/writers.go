package writers

import (
	"fmt"
	"io"
	"os"
)

type simpleWriter struct {
	name string
}

// 1 OMIT
func SimpleWriter(name string) *simpleWriter {
	return &simpleWriter{name}
}

func (w *simpleWriter) Write(b []byte) (int, error) {
	_, err := fmt.Fprintf(os.Stdout, "%q recorded %d bytes: %q\n", w.name, len(b), b)
	return len(b), err
}

// END 1 OMIT

// 2 OMIT
func (w *simpleWriter) ReadFrom(r io.Reader) (int64, error) {
	return io.Copy(w, r)
}

// END 2 OMIT
