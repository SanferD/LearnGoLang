package specialreader

import "io"

type SpecialReader struct {
	r io.Reader
	n int
}

func Create(r io.Reader, n int)  SpecialReader {
	return SpecialReader{r, n}
}

func (s SpecialReader) Read(p []byte) (int, error) {
	n, err := s.r.Read(p)
	
	if err != nil {
		return 0, err
	}
	
	if n > s.n {
		return s.n, io.EOF
	}

	return n, nil
}