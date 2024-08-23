package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestStos(t *testing.T) {
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

	input := []byte(`+
1
+
2
+
3
+
4
+
5
+
6
+
7
+
8
+
9
+
0
+
1
-
-
-
-
-
-
-
-
-
-
-`)

	wStdin.Write(input)
	wStdin.Close()

	stos()

	wStdout.Close()

	want := `:)
:)
:)
:)
:)
:)
:)
:)
:)
:)
:(
0
9
8
7
6
5
4
3
2
1
:(
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
