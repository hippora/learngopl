package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) == 0 {
		log.Fatal("parameter error")
	}
	for _, v := range os.Args[2:] {
		s := strings.Split(v, "=")
		go nc(s[0], s[1])
	}
	s := strings.Split(os.Args[1], "=")
	nc(s[0], s[1])
}

func nc(loc, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Printf("Location:%s\n", loc)
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
