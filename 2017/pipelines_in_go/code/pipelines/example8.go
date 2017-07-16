package main

import (
	"io"
	"os/exec"
	"github.com/msoap/byline"
	"regexp"
	"os"
	"compress/gzip"
)

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

func main() {
	cmdReader := CommandReader("find", ".", "-name", "*.go", "-type", "f")
	lineReader := byline.NewReader(cmdReader).GrepByRegexp(regexp.MustCompile("example.*"))
	teeReader := io.TeeReader(lineReader, os.Stdout)

	file, _ := os.Create("/tmp/gzipped_data")

	gzipWriter := gzip.NewWriter(file)
	io.Copy(gzipWriter, teeReader)
	gzipWriter.Close()
	file.Close()
	file.Close()
}
