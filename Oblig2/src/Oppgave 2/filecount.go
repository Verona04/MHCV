package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
		if len(os.Args) <= 1 {
		log.Fatalf("%s: missing file name\n", os.Args[0])
	}
	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%s: error reading file: %s: %v\n", os.Args[0], os.Args[1], err)
		return
	}
		defer f.Close()
		data := make([]byte, 100)

	newLines := 0

	var runeStats map[string]int
	runeStats = make(map[string]int)
	for {
		data = data[:cap(data)]
		n, err := f.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		data = data[:n]
			for _, b := range data {
			if _, ok := runeStats[string(b)]; !ok {
				runeStats[string(b)] = 1
			} else {
				runeStats[string(b)]++
			}
			if b == '\n' {
				newLines++
			}
		}
	}

	fmt.Printf("Information about %s:\n", filename)
	fmt.Printf("Number of lines in file:   %d\n", newLines)
	fmt.Printf("Most common runes:\n")

	n := map[int][]string{}
	var a []int
	for k, v := range runeStats {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	runeVal := 0

	for _, k := range a {
		for _, s := range n[k] {

			if s != " " && s != "\n" && s != "\r" && s != "\t" {
				runeVal++
				fmt.Printf("\t%d.  Rune: %s, Counts: %d\n", runeVal, s, k)
			}
		}
		if runeVal == 5 {
			break
		}
	}
}