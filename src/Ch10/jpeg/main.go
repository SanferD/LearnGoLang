package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
)

const (
	GIF = "GIF"
	JPG = "JPG"
	PNG = "PNG"
)

var destKind = flag.String("type", "jpeg", "the output types are " + strings.Join([]string{PNG, GIF, JPG}, " | "))

func main() {
	flag.Parse()
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return fmt.Errorf("toJPEG: error in decoding input: %s", err.Error())
	}
	fmt.Fprintf(os.Stderr, "kind = %s\n", kind)
	switch *destKind {
	case GIF:
		err = gif.Encode(out, img, nil)
	case PNG:
		err = png.Encode(out, img)
	case JPG:
		err = jpeg.Encode(out, img, nil)
	default:
		return fmt.Errorf("toJPEG: unknown output encoding: %s", *destKind)
	}
	if err != nil {
		return fmt.Errorf("toJPEG: could not encode: %s", err.Error())
	}
	return nil
}

