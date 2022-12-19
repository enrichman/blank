package main

import "io"

var _ io.Reader = (*MyReader)(nil)

type MyReader struct{}

func (r *MyReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

func main() {}
