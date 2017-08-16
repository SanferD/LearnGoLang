package main

import (
	"Ch7/7.2/specialwriter"
	"fmt"
	"io"
	"os"
)

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWriter := specialwriter.CreateSpecialWriter(w)
	return newWriter, newWriter.ByteCounter()
}

func main() {
	newStdout, counter := CountingWriter(os.Stdout)
	count, _ := fmt.Fprintf(newStdout, "abcd %s %s\n", "Heyo", "Technology")
	fmt.Println(count == int(*counter), count)
}