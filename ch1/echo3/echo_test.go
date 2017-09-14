package echo

import (
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	if false {
		fmt.Printf("AYY")
	}
}

func BenchmarkEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo2()
	}
}
