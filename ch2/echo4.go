// echo version 4
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Printf("sfsf\nsfsf\n")
	fmt.Printf(`sfsfsfs
		sfsfsfs
		sfsfsfs
	sfsf`)
}
