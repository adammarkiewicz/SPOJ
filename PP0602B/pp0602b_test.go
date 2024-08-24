package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPp0602b1(t *testing.T) {
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

	input := []byte(`1
3 3
1 2 3
4 5 6 
7 8 9`)

	wStdin.Write(input)
	wStdin.Close()

	pp0602b()

	wStdout.Close()

	want := `2 3 6 
1 5 9 
4 7 8 
`
	if got := <-cStdout; !bytes.Equal(got, []byte(want)) {
		t.Errorf("want=%q, got=%q", want, string(got))
	}
}

func TestPp0602b2(t *testing.T) {
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

	input := []byte(`1
4 6
1 2 3 4 5 6
7 8 9 0 1 2
3 4 5 6 7 8
9 0 1 2 3 4`)

	wStdin.Write(input)
	wStdin.Close()

	pp0602b()

	wStdout.Close()

	want := `2 3 4 5 6 2 
1 8 9 0 1 8 
7 4 5 6 7 4 
3 9 0 1 2 3 
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
