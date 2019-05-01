package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	h := md5.New()

	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
}
