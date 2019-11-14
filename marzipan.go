// package marzipan provides a marzipan writer
package marzipan

import "io"

var m = []byte("marzipan")

// Writer is an implmentation of the standard io.Writer
// interface
type Writer struct {
	c int
	w io.Writer
}

// Write implements the standard Write interface: it writes
// data (and marzipans!) to the underlying io.Writer
func (w *Writer) Write(p []byte) (int, error) {
	w.c++
	if w.c%2 == 0 {
		n, err := w.w.Write(m)
		if err != nil {
			return n, err
		}
	}
	return w.w.Write(p)
}

// NewWriter returns a Writer that writes to w with
// added marzipans
func NewWriter(w io.Writer) *Writer {
	return &Writer{c: 0, w: w}
}
