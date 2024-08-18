package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestSkarbfi(t *testing.T) {
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

	input := []byte(`3
3
1 1
0 2
3 1
4
0 1
2 1
1 1
3 1
2
0 1
0 2`)

	wStdin.Write(input)
	wStdin.Close()

	skarbfi()

	wStdout.Close()

	want := `0 1
3 1
studnia
0 3
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
