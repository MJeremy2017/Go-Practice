package main

import (
	"fmt"
	"testing"
)

func BenchmarkSprinf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}
