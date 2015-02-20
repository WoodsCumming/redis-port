// Copyright 2014 Wandoujia Inc. All Rights Reserved.
// Licensed under the MIT (MIT-LICENSE.txt) license.

package ioutils

import (
	"io"

	"github.com/wandoulabs/redis-port/pkg/libs/counter"
)

type CountReader struct {
	p *counter.Int64
	r io.Reader
}

func NewCountReader(r io.Reader, p *counter.Int64) *CountReader {
	if p == nil {
		p = &counter.Int64{}
	}
	return &CountReader{p: p, r: r}
}

func (r *CountReader) Count() int64 {
	return r.p.Get()
}

func (r *CountReader) ResetCounter() int64 {
	return r.p.Reset()
}

func (r *CountReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	r.p.Add(int64(n))
	return n, err
}

type CountWriter struct {
	p *counter.Int64
	w io.Writer
}

func NewCountWriter(w io.Writer, p *counter.Int64) *CountWriter {
	if p == nil {
		p = &counter.Int64{}
	}
	return &CountWriter{p: p, w: w}
}

func (w *CountWriter) Count() int64 {
	return w.p.Get()
}

func (w *CountWriter) ResetCounter() int64 {
	return w.p.Reset()
}

func (w *CountWriter) Write(p []byte) (int, error) {
	n, err := w.w.Write(p)
	w.p.Add(int64(n))
	return n, err
}
