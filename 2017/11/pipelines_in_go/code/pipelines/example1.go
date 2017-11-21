package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/superstas/talks/2017/10/pipelines_in_go/code/pipelines/utils"
)

func init() {
	// ugly hack to execute STDIN example in go present :(
	os.Stdin = utils.MakeTmpFile("/tmp/1")
}

func main() {
	// 1 OMIT
	r := io.LimitReader(os.Stdin, 8) // HL
	res, err := ioutil.ReadAll(r)
	fmt.Printf("res: %s\nerr: %v\n", res, err)
	// END 1 OMIT
}
