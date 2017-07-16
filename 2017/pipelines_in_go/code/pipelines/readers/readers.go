package readers

import "io"


type simpleReader struct {
	data []byte
}
// 1 OMIT
func SimpleReader(s string) *simpleReader {
	return &simpleReader{[]byte(s)}
}

func (r *simpleReader) Read(p []byte) (int, error) {
	n := copy(p, r.data) // HL
	if n == 0 {
		return 0, io.ErrUnexpectedEOF
	}
	r.data = r.data[:0] // HL
	return n, io.EOF
}
// END 1 OMIT

// 3 OMIT
func (r *simpleReader) WriteTo(w io.Writer) (int, error) {
	return w.Write(r.data)
}
// END 3 OMIT


type brokenReader struct {
	io.Reader
}

// 2 OMIT
func BrokenReader(r io.Reader) io.Reader {
	return &brokenReader{r}
}

func (r *brokenReader) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	if err != nil && err != io.EOF {
		return 0, err
	}
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			continue
		}
		p[i] = byte('#') // HL
	}
	return n, err
}
// END 2 OMIT
