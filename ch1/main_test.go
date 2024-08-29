package main

import (
	"os"
	"strings"
	"testing"
)

// Benchmark for inefficient string concatenation
func BenchmarkInefficientConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s, sep string
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
	}
}

// Benchmark for efficient concatenation using strings.Join
func BenchmarkEfficientJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(os.Args[1:], " ")
	}
}
