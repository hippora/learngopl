// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Println(os.Getwd())
}
