package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestCalc2(t *testing.T) {
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

	input := []byte(`z 3 6
z 1 89
z 2 60
z 0 11
+ 0 1
- 1 2
* 2 3
/ 3 0
% 3 1`)

	wStdin.Write(input)
	wStdin.Close()

	calc2()

	wStdout.Close()

	want := `100
29
360
0
6
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
