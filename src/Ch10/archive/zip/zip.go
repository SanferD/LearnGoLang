package zip

import (
	"image"
	"io"
)

func Decode(r io.Reader) (image.Image, error) {
	return nil, nil
}

func DecodeConfig(r io.Reader) (image.Config, error) {
	return image.Config{}, nil
}

func init() {
	header := string([]byte{0x50, 0x4B, 0x03, 0x04})
	image.RegisterFormat("zip", header, Decode, DecodeConfig)
}
