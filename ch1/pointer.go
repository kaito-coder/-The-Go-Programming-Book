package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing new line")
var s = flag.String("s", " ", "separator")
func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*s))
	if !*n {
		fmt.Println()
	}
}
func incr(p *int)int  {
	*p ++
	return *p
}