// Package limit provides a way to limit the amount of bytes read from an
// io.ReadCloser.
package limit

import "io"

// LimitedReadCloser reads from R but limits the amount of data returned to just
// N bytes. Each call to Read updates N to reflect the new amount remaining.
// Read returns EOF when N <= 0 or when the underlying R returns EOF or another
// error.
type LimitedReadCloser struct {
	R io.ReadCloser // underlying reader
	N int64         // max bytes remaining
}

// ReadCloser returns an io.ReadCloser that reads from r but stops with EOF after n
// bytes. The underlying implementation is a *LimitedReadCloser.
func ReadCloser(r io.ReadCloser, n int64) io.ReadCloser {
	return &LimitedReadCloser{R: r, N: n}
}

func (l *LimitedReadCloser) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}

	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

// Close closes the underlying reader R.
func (l *LimitedReadCloser) Close() error {
	return l.R.Close()
}
