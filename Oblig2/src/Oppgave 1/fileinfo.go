package main

import (
	"fmt"
	"os"
	"log"
)

func main() {

	fi, err := os.Lstat("fileinfo.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name: ", fi.Name())
	fmt.Println("Size : ", fi.Size())
	switch mode := fi.Mode(); {
	case mode.IsDir():
		fmt.Println("directory")
	case mode.IsRegular():
		fmt.Println("Regular file")
		fmt.Println("UnixPermissionBits", mode.Perm())
	case os.ModeAppend != 0:
		fmt.Println("Is Append")
	case os.ModeDevice != 0:
		fmt.Println("Is a Device file")
	case os.ModeSymlink != 0:
		fmt.Println("Is a Symbolic Link")
	}
}
