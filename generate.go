// Copyright 2017 The OLX Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"os"
	"fmt"
	"flag"
)

var data = []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus quis dolor non nisl lobortis dictum id vitae metus. Quisque pharetra neque quis metus congue laoreet. In finibus gravida lectus id mollis. Ut bibendum neque in magna fringilla volutpat. Fusce tempor posuere quam, vitae vehicula nisi dapibus placerat. Nam sed maximus diam, non faucibus elit. Nunc ultricies dui nec pellentesque pellentesque. Ut augue dui, condimentum quis vestibulum et, egestas sit amet nulla. Nam elementum mauris nec quam molestie, efficitur tempor metus dapibus. In vestibulum pretium dolor et tristique. Phasellus eu magna dignissim, ornare ex id, fringilla mauris. Pellentesque sed orci ligula. Donec ante neque, euismod quis erat eget, suscipit bibendum erat.\n")
var fileName string = "data.json"
var fileSize int = 1024 * 64

func init() {
	flag.StringVar(&fileName, "fileName", "data.json", "Define the name of json")
	flag.IntVar(&fileSize, "bytes", 12558918857, "Define the number of bytes (default 12558918857) (12GB).")
	flag.Parse()
}

func main() {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error to create [%s]: %s", fileName, err)
		os.Exit(1)
	}
	defer file.Close()

	var tmpSize int = 0
	var dataSize int = len(data)
	for {
		next := tmpSize + dataSize

		// if file size will be more then fileSize it will adjust
		if next > fileSize {
			next = dataSize - (next - fileSize)
			data = data[:next]
		}

		file.Write(data) // write chunk
		tmpSize = next   // update tmpSize

		if tmpSize >= fileSize {
			break
		}
	}

	fmt.Printf("tmp: %d - size: %d \n", tmpSize, fileSize)
}
