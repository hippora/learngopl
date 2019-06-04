package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	val := flag.String("t", "sha256", "option:sha256 ,sha384, sha512")
	flag.Parse()
	if len(flag.Arg(0)) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println(*val, "-", flag.Arg(0))
	switch *val {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(flag.Arg(0))))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(flag.Arg(0))))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(flag.Arg(0))))
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
