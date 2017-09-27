package readers

import "syscall"

// 1 OMIT
var Stdin = NewFile(uintptr(syscall.Stdin), "/dev/stdin")

// END 1 OMIT
