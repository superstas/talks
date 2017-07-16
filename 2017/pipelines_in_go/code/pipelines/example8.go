package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/msoap/byline"
)

// 1 OMIT
func CommandReader(cmd string, arg ...string) *commandReader {
	return &commandReader{cmd: exec.Command(cmd, arg...)}
}

type commandReader struct {
	cmd *exec.Cmd
}

func (r *commandReader) Read(p []byte) (int, error) {
	reader, err := r.cmd.StdoutPipe()
	if err != nil {
		return 0, io.EOF
	}
	r.cmd.Start()
	defer r.cmd.Wait()
	return reader.Read(p)
}

// END 1 OMIT

func main() {
	// 2 OMIT
	cmdReader := CommandReader("find", ".", "-name", "*.go", "-type", "f")
	// END 2 OMIT

	// 3 OMIT
	lineReader := byline.NewReader(cmdReader).GrepByRegexp(regexp.MustCompile("example.*"))
	// END 3 OMIT

	// 4 OMIT
	teeReader := io.TeeReader(lineReader, os.Stdout)
	// END 4 OMIT

	// 5 OMIT
	file, err := os.Create("/tmp/gzipped_data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// END 5 OMIT

	// 6 OMIT
	gzipWriter := gzip.NewWriter(file)
	io.Copy(gzipWriter, teeReader) // HL
	gzipWriter.Close()
	// END 6 OMIT
}
