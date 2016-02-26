package got

import "io"

// A simple limitedReader similar to io.LimitReader that also let's us know
// if we reached EOF
type limitedReader struct {
	reader    io.Reader
	maxBytes  int64
	lastError error
}

func (l *limitedReader) Read(p []byte) (int, error) {
	if l.maxBytes <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.maxBytes {
		p = p[0:l.maxBytes]
	}
	n, err := l.reader.Read(p)
	l.lastError = err
	l.maxBytes -= int64(n)
	return n, err
}

func (l *limitedReader) ReachedEOF() bool {
	return l.lastError == io.EOF
}
