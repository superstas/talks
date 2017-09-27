package readers

import (
	"io"
)

type simpleReader struct {
	data []byte
}

// 1 OMIT
func SimpleReader(s string) *simpleReader {
	return &simpleReader{[]byte(s)}
}

func (r *simpleReader) Read(p []byte) (int, error) {
	n := copy(p, r.data) // HL
	// todo: partial read support
	r.data = r.data[n:] // HL
	return n, io.EOF    // read finished
}

// END 1 OMIT

// 3 OMIT
func (r *simpleReader) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(r.data) // HL
	r.data = r.data[n:]
	return int64(n), err
}

// END 3 OMIT

type brokenReader struct {
	io.Reader
}

// 2 OMIT
func BrokenReader(r io.Reader) io.Reader {
	return &brokenReader{r}
}

// END 2 OMIT

// 2_1 OMIT
func (r *brokenReader) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	// todo: check errors, add partial read support
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			p[i] = byte('#') // HL
		}
	}
	return n, err
}

// END 2_1 OMIT
