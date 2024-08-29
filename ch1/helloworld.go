package main

import (
	"fmt"
	"os"
	"testing"
	"strings"
)

func main() {

	//echo 1
	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Println(os.Args[i])
	// 	s += sep + os.Args[i]
	// }

	//echo 2
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	// echo 3
//	fmt.Println(os.Args[1:])

}
func concat(arg []string) string {
	s, sep := "", ""
	for _, arg := range arg[1:] {
		fmt.Println(arg)
		s += sep + arg
		sep = " "
	}
	return s
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat(os.Args[1:])
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(os.Args[1:], " ")
	}
}