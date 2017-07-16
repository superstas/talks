package writers

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func SimpleWriter(name string) *simpleWriter {
	return &simpleWriter{name}
}

type simpleWriter struct {
	name string
}

// 1 OMIT
func (w *simpleWriter) Write(b []byte) (int, error) {
	return fmt.Fprintf(os.Stdout, "%q wrote %d bytes: %v\n", w.name, len(b), b)
}

// END 1 OMIT

// 2 OMIT
func (w *simpleWriter) ReadFrom(r io.Reader) (int, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return 0, err
	}
	return fmt.Fprintf(os.Stdout, "%q wrote %d bytes via ReadFrom: %v\n", w.name, len(b), b)
}

// END 2 OMIT
