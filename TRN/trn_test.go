package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestTrn(t *testing.T) {
	rStdin, wStdin, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	rStdout, wStdout, _ := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	cStdout := make(chan []byte)
	go readAll(t, rStdout, cStdout)

	defer func(f *os.File) { os.Stdin = f }(os.Stdin)
	defer func(f *os.File) { os.Stdout = f }(os.Stdout)
	os.Stdin = rStdin
	os.Stdout = wStdout

	input := []byte(`4 3
1 2 5
4 3 3
3 4 9
8 7 7`)

	wStdin.Write(input)
	wStdin.Close()

	trn()

	wStdout.Close()

	want := `1 4 3 8 
2 3 4 7 
5 3 9 7 
`
	if got := <-cStdout; !bytes.Equal(got, []byte(want)) {
		t.Errorf("want=%q, got=%q", want, string(got))
	}
}

func readAll(t *testing.T, r io.Reader, c chan<- []byte) {
	data, err := io.ReadAll(r)
	if err != nil {
		t.Error(err)
	}
	c <- data
}
