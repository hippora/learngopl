package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string = "test"

func main() {
	fmt.Printf("%b\n", byte(255))
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%s", cwd)
	p()
}

func p() {
	log.Printf("%s", cwd)
}
