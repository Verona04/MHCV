package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"os"
	"log"
	"fmt"
)

func main() {
	readFromFile("numbers.txt")
}

func readFromFile(filePath string){

	read, err := ioutil.ReadFile(filePath)

	content := string(read)

	tempContent := strings.Split(content, "\n")
	string1, string2 := tempContent[0], tempContent[1]

	int1, err := strconv.Atoi(string1)
	int2, err := strconv.Atoi(string2)

	result := (int1+int2)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintln(result))
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}


}