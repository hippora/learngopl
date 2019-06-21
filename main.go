package main

import (
	//. "fmt"
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	for i := 0; i < 50; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}

}
