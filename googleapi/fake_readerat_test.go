// Copyright (C) 2015 Motorola Mobility LLC All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package googleapi

import (
	"bytes"
	"io"
	"testing"
)

func setup100() *fakeReaderAt {
	data := make([]byte, 100)
	for i := byte(0); i < 100; i++ {
		data[i] = i
	}
	r := bytes.NewReader(data)
	return &fakeReaderAt{r: r}
}

func expect(t *testing.T, fr *fakeReaderAt, buf []byte, off int64, n int, err error, first, last byte, msg string) {
	n1, err1 := fr.ReadAt(buf, off)
	if n != n1 {
		t.Fatalf("%s: expected %v byte got %v, err %v", msg, n, n1, err1)
	} else if err1 != err {
		t.Fatalf("%s: expected err %v got %v", msg, err, err1)
	} else if buf[0] != first || buf[len(buf)-1] != last {
		t.Fatalf("%s: wrong buf content, expected [%v...%v] got %v", msg, first, last, buf)
	}
}

func expectSize(t *testing.T, fr *fakeReaderAt, off int64, size int, n int, err error, msg string) {
	n1, err1 := fr.SizeAt(off, size)
	if n != n1 {
		t.Fatalf("%s: expected %v size got %v, err %v", msg, n, n1, err1)
	} else if err1 != err {
		t.Fatalf("%s: expected err %v got %v", msg, err, err1)
	}
}

func TestReadSkipRead(t *testing.T) {
	fr := setup100()
	buf := make([]byte, 10)
	expect(t, fr, buf, 0, 10, nil, 0, 9, "initial read")
	expect(t, fr, buf, 20, 10, nil, 20, 29, "skip read 1")
	expect(t, fr, buf, 40, 10, nil, 40, 49, "skip read 2")
}

func TestSkipInit(t *testing.T) {
	fr := setup100()
	buf := make([]byte, 10)
	expect(t, fr, buf, 20, 10, nil, 20, 29, "skip read")
}

func TestBacktrack(t *testing.T) {
	fr := setup100()
	buf := make([]byte, 10)
	expect(t, fr, buf, 10, 10, nil, 10, 19, "initial read")
	expect(t, fr, buf[0:5], 13, 5, nil, 13, 17, "back tracked read")
	expect(t, fr, buf, 15, 10, nil, 15, 24, "back tracked + new read")
	expect(t, fr, buf[0:5], 20, 5, nil, 20, 24, "back tracked to the end read")
}

func TestSKipToENd(t *testing.T) {
	fr := setup100()
	buf := make([]byte, 10)
	expect(t, fr, buf, 0, 10, nil, 0, 9, "initial read")
	expect(t, fr, buf, 90, 10, io.EOF, 90, 99, "end read")
}

func TestBacktrackErr(t *testing.T) {
	fr := setup100()
	buf := make([]byte, 10)
	expect(t, fr, buf, 20, 10, nil, 20, 29, "initial read")
	expect(t, fr, buf, 19, 0, errBacktrack, buf[0], buf[9], "error read")
}

func TestEOFErr(t *testing.T) {
	fr := setup100()
	buf := make([]byte, 10)
	expect(t, fr, buf, 95, 5, io.EOF, 95, buf[9], "initial read")
}

func TestEOFErr2(t *testing.T) {
	fr := setup100()
	expectSize(t, fr, 10, 10, 10, nil, "first sizing")
	expectSize(t, fr, 110, 10, 0, io.EOF, "skip past end")
}

func TestSizeAt(t *testing.T) {
	fr := setup100()
	buf := make([]byte, 10)
	expectSize(t, fr, 10, 20, 20, nil, "first sizing")
	expectSize(t, fr, 10, 5, 5, nil, "second sizing")
	expectSize(t, fr, 10, 25, 25, nil, "third sizing")
	expect(t, fr, buf, 20, 10, nil, 20, 29, "read")
	expectSize(t, fr, 80, 20, 20, io.EOF, "size when aligned to EOF")
	expectSize(t, fr, 90, 20, 10, io.EOF, "size when passed EOF")
}
