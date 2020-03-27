package helpers

import (
	"fmt"
	"io"
	"os"
)

var writer io.Writer

func PrintInit(out io.Writer) {
	writer = out
}

func Println(msg ...interface{}) {
	check()
	fmt.Fprintln(writer, msg...)
}

func Printf(msg string, args ...interface{}) {
	check()
	fmt.Fprintf(writer, msg, args...)
}

func check() {
	if writer == nil {
		writer = os.Stdout
	}
}
