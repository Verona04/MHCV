package main

import (
	"fmt"
	"os"
	"log"
)

func main() {

	if len(os.Args) <= 1 {
		log.Fatalf("%s: missing file name\n", os.Args[0])
	}
	filename := os.Args[1]

	fi, err := os.Lstat(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Information about file " + filename + ":")
	fiBytes := fi.Size()
	fiKB := float64(fi.Size()) / 1024
	fiMB := float64(fi.Size()) / 1024 / 1024
	fiGB := float64(fi.Size()) / 1024 / 1024 / 1024
	fmt.Printf("Size: %d bytes, %f KB, %f MB and %f GB\n", fiBytes, fiKB, fiMB, fiGB)
	if fi.IsDir() {
		fmt.Println("Is a directory")
	} else {
	fmt.Println("Is not a directory")
	}
	mode := fi.Mode()
	if mode.IsRegular() {
		fmt.Println("Is a regular file")
	} else {
		fmt.Println("Is not a regular file")
	}
	perm := mode.Perm()
	fmt.Printf("Has Unix permission bits: %s\n", perm)

	if mode&os.ModeAppend != 0 {
		fmt.Println("Is append only")
	} else {
	fmt.Println("Is not append only")
	}
	if mode&os.ModeDevice != 0 {
		fmt.Println("Is a device file")
	} else {
		fmt.Println("Is not a device file")
	}
	if mode&os.ModeSymlink != 0 {
		fmt.Println("Is a symbolic Link")
	} else {
		fmt.Println("Is not a symbolic Link")
	}
}
