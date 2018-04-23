package main

import (
	"fmt"
//	"time"
	"log"
	"os"
	"os/signal"
	"syscall"
)


func main() {
	c := make(chan int)
	inputNumbers(c)
}


func inputNumbers(c chan int){
	var a, b int
	h := make(chan os.Signal, 2)
	signal.Notify(h, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-h
		log.Fatal("Exiting")
		os.Exit(1)
	}()
	fmt.Println("Enter a number: ")

	_, err := fmt.Scanf("%d", &a)

	if err != nil {
		log.Fatal("Unable to read number: ", err)
	}
	fmt.Println("Enter another number: ")

	fmt.Scanf("%d", &b)

	if err != nil {
		log.Fatal("Unable to read number: ", err)
	}
	fmt.Println("You entered the numbers ", a, "and ", b)
	go func(a int, b int) { //fikk issues med fatal error: all goroutines are asleep - deadlock, så valgte å bruke en anonym func
		c <- a
		c <- b  }(a,b)

	addNumbers(c)

	res := <- c
	fmt.Println("Sum numbers: ", res)

}

func addNumbers(c chan int){
	a, b := <-c, <-c
	res := (a+b)

	go func() { c <- res }()
}