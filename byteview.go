package icache

import (
	"bytes"
	"errors"
	"io"
	"strings"
)

// A ByteView holds an immutable view of bytes.
// Internally it wraps either a []byte or a string,
// but that details is invisible to callers.
//
// The byte was chosen to support storage of arbitrary
// data types, such as strings, images, etc.
type ByteView struct {
	// If b is non-nil, b is used, else s is used.
	b []byte
	s string
}

// Len returns the length of the view.
func (v ByteView) Len() int {
	if v.b != nil {
		return len(v.b)
	}
	return len(v.s)
}

// ByteSlice returns a copy of the data as a byte slice.
func (v ByteView) ByteSlice() []byte {
	if v.b != nil {
		return cloneBytes(v.b)
	}
	return []byte(v.s)
}

// String returns the data as a string.
func (v ByteView) String() string {
	if v.b != nil {
		return string(v.b)
	}
	return v.s
}

// At returns the byte at index i.
func (v ByteView) At(i int) byte {
	if v.b != nil {
		return v.b[i]
	}
	return v.s[i]
}

// Slice slices the view between the provided from and to indices.
func (v ByteView) Slice(from, to int) ByteView {
	if v.b != nil {
		return ByteView{b: v.b[from:to]}
	}
	return ByteView{s: v.s[from:to]}
}

// SliceFrom slices the view from the provided index until the end.
func (v ByteView) SliceFrom(from int) ByteView {
	if v.b != nil {
		return ByteView{b: v.b[from:]}
	}
	return ByteView{s: v.s[from:]}
}

// Copy copies b into dest and returns the number of bytes copied.
func (v ByteView) Copy(dest []byte) int {
	if v.b != nil {
		return copy(dest, v.b)
	}
	return copy(dest, v.s)
}

// Equal returns whether the bytes in b are the same as
// the bytes in b2.
func (v ByteView) Equal(b2 ByteView) bool {
	if b2.b == nil {
		return v.EqualString(b2.s)
	}
	return v.EqualBytes(b2.b)
}

// EqualString returns whether the bytes in b are the same
// as the bytes in s.
func (v ByteView) EqualString(s string) bool {
	if v.b == nil {
		return v.s == s
	}

	l := v.Len()
	if l != len(s) {
		return false
	}
	for i, bi := range v.b {
		if bi != s[i] {
			return false
		}
	}
	return true
}

// EqualBytes returns whether the bytes in b are the same
// as the bytes in b2.
func (v ByteView) EqualBytes(b2 []byte) bool {
	if v.b != nil {
		return bytes.Equal(v.b, b2)
	}

	l := v.Len()
	if l != len(b2) {
		return false
	}
	for i, bi := range b2 {
		if bi != v.s[i] {
			return false
		}
	}
	return true
}

// Reader returns an io.ReadSeeker for the bytes in v.
func (v ByteView) Reader() io.ReadSeeker {
	if v.b != nil {
		return bytes.NewReader(v.b)
	}
	return strings.NewReader(v.s)
}

// ReadAt implements io.ReaderAt on the bytes in v.
func (v ByteView) ReadAt(p []byte, off int64) (n int, err error) {
	if off < 0 {
		return 0, errors.New("view: invalid offset")
	}
	if off >= int64(v.Len()) {
		return 0, io.EOF
	}
	n = v.SliceFrom(int(off)).Copy(p)
	if n < len(p) {
		err = io.EOF
	}
	return
}

// WriteTo implements io.WriterTo on the bytes in v.
func (v ByteView) WriteTo(w io.Writer) (n int64, err error) {
	var m int
	if v.b != nil {
		m, err = w.Write(v.b)
	} else {
		m, err = io.WriteString(w, v.s)
	}
	if err == nil && m < v.Len() {
		err = io.ErrShortWrite
	}
	n = int64(m)
	return
}