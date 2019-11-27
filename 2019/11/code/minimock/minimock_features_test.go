package code

import (
	"fmt"
	"io"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

// 1 OMIT
type CustomReader struct {
	io.ReadCloser
}

func (r CustomReader) Read(p []byte) (int, error) {
	return r.ReadCloser.Read(p)
}

// END 1 OMIT

func Test_Expect_Return_Example(t *testing.T) {
	// 2 OMIT
	readCloserMock := NewReadCloserMock(t).
		ReadMock.Expect([]byte{0, 0, 0}).Return(3, nil). // HL
		CloseMock.Return(nil)                            // HL

	reader := CustomReader{readCloserMock}

	n, err := reader.Read(make([]byte, 3)) // HL
	// ReadCloserMock.Read got unexpected parameters, want ...

	reader.Close() // HL
	// Unexpected call to ReadCloserMock.Close ...
	// END 2 OMIT

	assert.NoError(t, err)
	assert.Equal(t, 3, n)
}

func Test_Expect_Return_Example_2(t *testing.T) {
	// 3 OMIT
	readCloserMock := NewReadCloserMock(t).ReadMock.Inspect(func(p []byte) {
		assert.EqualValues(t, 0, p[1]) // HL
	}).Return(3, nil)
	// END 3 OMIT
	reader := CustomReader{readCloserMock}

	n, err := reader.Read(make([]byte, 3))

	assert.NoError(t, err)
	assert.Equal(t, 3, n)
}

func Test_When_Then_Example(t *testing.T) {
	// 4 OMIT
	readCloserMock := NewReadCloserMock(t).
		ReadMock.When([]byte{}).Then(0, io.EOF).    // HL
		ReadMock.When([]byte{0}).Then(1, nil).      // HL
		ReadMock.When([]byte{0, 0, 0}).Then(3, nil) // HL

	reader := CustomReader{readCloserMock}

	n, err := reader.Read([]byte{})
	// err == io.EOF
	// END 4 OMIT
	assert.Error(t, io.EOF, err.Error())
	assert.Equal(t, 0, n)
}

func Test_Finish_Example(t *testing.T) {
	// 5 OMIT
	mc := minimock.NewController(t) // HL
	defer mc.Finish()               // HL

	stringerMock := NewStringerMock(mc)
	stringerMock.StringMock.Return("minimock")
	// Expected call to StringerMock.String // HL

	readCloserMock := NewReadCloserMock(mc)
	readCloserMock.ReadMock.Return(5, nil)
	// END 5 OMIT
	fmt.Print(stringerMock)
	readCloserMock.Read([]byte{})
}
