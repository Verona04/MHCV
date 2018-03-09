package main

import (
	"fmt"
	//	"time"
	"os"
	"log"
)


func main() {
	inputToFile()
}


func inputToFile(){
	var a, b int
	fmt.Println("Enter a number: ")

	_, err := fmt.Scanf("%d", &a)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Enter another number: ")

	fmt.Scanf("%d", &b)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("You entered the numbers ", a, "and ", b)

	file, err := os.Create("numbers.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintln(a))
	_, err = file.WriteString(fmt.Sprintln(b))

}