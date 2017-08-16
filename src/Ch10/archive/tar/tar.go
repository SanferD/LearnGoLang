package tar

import (
	"io"
	"image"
	"strings"
)

func Decode(r io.Reader) (image.Image, error) {
	return nil, nil
}

func DecodeConfig(r io.Reader) (image.Config, error) {
	return image.Config{}, nil
}

func init() {
	offset := strings.Repeat("?", 0x101)
	header := string([]byte{0x75, 0x73, 0x74, 0x61, 0x72, 0x00, 0x30, 0x30})
	header = offset + header
	image.RegisterFormat("tar", header, Decode, DecodeConfig)
}
