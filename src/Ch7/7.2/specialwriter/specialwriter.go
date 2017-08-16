package specialwriter

import "io"

type SpecialWriter struct {
	byteCounter *int64
	w io.Writer
}

func (m SpecialWriter) Write(p []byte) (int, error) {
	*m.byteCounter += int64(len(p))
	return m.w.Write(p) 
}

func (m SpecialWriter) ByteCounter() *int64 {
	return m.byteCounter
}

func CreateSpecialWriter(w io.Writer) SpecialWriter {
	return SpecialWriter{ new(int64), w }
}
