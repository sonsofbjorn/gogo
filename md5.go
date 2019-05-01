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

	// Sum(data []byte) returna a slice of [Size]byte
	for _, s := range h.Sum(nil) {
		fmt.Printf("%x\n", s)
	}

	md5 := h.Sum(nil)

	// %x base16 lower-case, %X base16 upper-case
	fmt.Printf("%X%x\n", md5[0:2], md5[2:4])
}
