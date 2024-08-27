package main

import (
    "bytes"
    "io"
    "log"
    "os"
    "testing"
)

func TestProblem(t *testing.T) {
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

    wStdin.Write(inputTestData())
    wStdin.Close()

    main()

    wStdout.Close()

    want := outputTestData()
    if got := <-cStdout; !bytes.Equal(got, []byte(want)) {
        t.Errorf("want=%q, got=%q", want, string(got))
    }
}

func readAll(t *testing.T, r io.Reader, c chan<- []byte) {
    data, err := io.ReadAll(r)
    if err != nil {
        t.Error(err)
    }

    if len(data) > 0 && data[len(data)-1] == '\n' {
        data = data[:len(data)-1]
    }

    c <- data
}

func inputTestData() []byte {
    data, err := os.ReadFile("test.input")
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }

    return data
}

func outputTestData() []byte {
    data, err := os.ReadFile("test.output")
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }

    data = bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))

    return data
}
