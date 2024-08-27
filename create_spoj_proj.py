#!/usr/bin/env python3

import os, sys, re, time, requests, itertools, threading, subprocess
from bs4 import BeautifulSoup
from markdownify import markdownify as md

PROBLEM_NAME_REGEX_PATTERN = r"^[A-Z0-9_]+$"

SPOJ_PROBLEM_BASE_URL = "https://pl.spoj.com/problems/"

README_FILENAME = "README.md"
README_ENCODING = "utf-8"

TEST_INPUT_FILENAME = "test.input"
TEST_OUTPUT_FILENAME = "test.output"

GO_PROGRAM_FILENAME = "main.go"
GO_PROGRAM_TEMPLATE = """package main
                         
func main() {

}"""

GO_PROBLEM_TEST_FILENAME = "problem_test.go"
GO_PROBLEM_TEST = """package main

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

    if len(data) > 0 && data[len(data)-1] == '\\n' {
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

    data = bytes.ReplaceAll(data, []byte("\\r\\n"), []byte("\\n"))

    return data
}
"""


def create_go_project(problem):
    if not re.match(PROBLEM_NAME_REGEX_PATTERN, problem):
        print(f"Problem name '{problem}' is not valid.")
        sys.exit()

    if os.path.isdir(problem):
        print(f"Directory for problem '{problem}' already exists.")
        sys.exit()

    try:
        response = requests.get(SPOJ_PROBLEM_BASE_URL + problem)
        if b"In a few seconds you will be redirected to" in response.content:
            print(f"Problem '{problem}' does not exist in SPOJ.")
            sys.exit()
    except requests.exceptions.RequestException as e:
        print(f"Error while getting the page '{SPOJ_PROBLEM_BASE_URL + problem}': {e}")
        sys.exit()

    os.mkdir(problem)

    soup = BeautifulSoup(response.content, 'html.parser')

    ################### go.mod ###################
    command = f"go mod init {problem}"
    working_directory = problem

    result = subprocess.run(command, shell=True, cwd=working_directory, text=True, capture_output=True)
    
    print(result.stdout)
    print(result.stderr)
    
    ################### README.md ###################
    print(f"create_spoj_proj: creating {problem}\\README.md")

    divs = soup.find_all('div', {'class': None, 'id': None, 'style': None})
    readme = md(divs[1].text).replace('\\', '')
    readme_file = problem + '\\' + README_FILENAME
    with open(readme_file, 'w', encoding=README_ENCODING) as file:
        file.write(readme)

    ################### test.input ###################
    print(f"create_spoj_proj: creating {problem}\\test.input")

    input = ""
    strong_inputs = soup.find_all("strong", string="Input:")
    for strong_input in strong_inputs:
        for sibling in strong_input.next_siblings:
            if "Output" in sibling.text:
                break
            input += sibling.text.strip() + '\n'

    input = input.replace("\r", "")
    input = input.strip()

    test_input_file = problem + '\\' + TEST_INPUT_FILENAME
    with open(test_input_file, 'w') as file:
        file.write(input)

    ################### test.output ###################
    print(f"create_spoj_proj: creating {problem}\\test.output")

    output = ""
    strong_outputs = soup.find_all("strong", string="Output:")
    for strong_output in strong_outputs:
        for sibling in strong_output.next_siblings:
            output += sibling.text.strip() + '\n'

    output = output.replace("\r", "")
    output = output.strip()

    test_output_file = problem + '\\' + TEST_OUTPUT_FILENAME
    with open(test_output_file, 'w') as file:
        file.write(output)

    ################### main.go ###################
    print(f"create_spoj_proj: creating {problem}\\main.go")

    program_file = problem + '\\' + GO_PROGRAM_FILENAME
    with open(program_file, 'w') as file:
        file.write(GO_PROGRAM_TEMPLATE)

    ################### problem_test.go ###################
    print(f"create_spoj_proj: creating {problem}\\problem_test.go")

    problem_test_file = problem + '\\' + GO_PROBLEM_TEST_FILENAME
    with open(problem_test_file, 'w') as file:
        file.write(GO_PROBLEM_TEST)


problem = sys.argv[1]
create_go_project(problem)
