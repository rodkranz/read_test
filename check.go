// Copyright 2017 The OLX Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := "unknown"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	text := "Started Reading json file."
	if !strings.Contains(string(line), text) {
		fmt.Printf("Text is missing: %s\n", text)
		os.Exit(1)
	}

	line, _, err = reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error to open data.json")
		os.Exit(1)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println("Error to open data.json")
		os.Exit(1)
	}

	const chunk = 64 * 1024
	total := (fileInfo.Size() / chunk) + 1
	text = fmt.Sprintf("Bytes: %d Chunks: %d", fileInfo.Size(), total)
	if string(line) == text {
		fmt.Printf("The results are correct for [%s].\n", name)
		os.Exit(0)
	}

	fmt.Printf("The results are wrong for [%s].\n", name)
	fmt.Printf("\t Expected: %s\n", text)
	fmt.Printf("\t   Actual: %s\n", line)
	os.Exit(1)
}
