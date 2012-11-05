package enmime

import (
	"io"
)

// Base64Cleaner helps work around bugs in Go's built-in base64 decoder
// by stripping out whitespace that would cause Go to lose count of
// things and issue an "illegal base64 data at input byte..." error
type Base64Cleaner struct {
	in    io.Reader
	buf   [1024]byte
	//count int64
}

func NewBase64Cleaner(r io.Reader) *Base64Cleaner {
	return &Base64Cleaner{in: r}
}

func (qp *Base64Cleaner) Read(p []byte) (n int, err error) {
	// Size our slice to theirs
	size := len(qp.buf)
	if len(p) < size {
		size = len(p)
	}
	buf := qp.buf[:size]
	bn, err := qp.in.Read(buf)
	for i := 0; i < bn; i++ {
		switch buf[i] {
		case ' ', '\t', '\r', '\n':
			// Strip these
		default:
			p[n] = buf[i]
			n++
		}
	}
	// Count may be useful if I need to pad to even quads
	//qp.count += int64(n)
	return n, err
}
