package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	const filename = "/tmp/file.txt"

	err := ioutil.WriteFile(filename, []byte("Hello, file system\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)
}

// When a process starts, the file system is populated with some devices under /dev and an empty /tmp directory. The program can manipulate the file system as usual, but when the process exits any changes to the file system are lost.
