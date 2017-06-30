// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "../data.json"
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error to read [file=%v]: %v\n", fileName, err.Error())
		os.Exit(1)
	}

	fmt.Println("Started Reading json file.")

	nBytes, nChunks := int64(0), int64(0)
	r := bufio.NewReader(f)
	buf := make([]byte, 0, 64*1024)
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		nChunks++
		nBytes += int64(len(buf))
		// process buf
		if err != nil && err != io.EOF {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("Bytes:", nBytes, "Chunks:", nChunks)
	os.Exit(0)
}
