package debuglog_test

import (
	"io"
	"log"
	"testing"

	"github.com/paskozdilar/debuglog"
)

// Benchmark loop with debuglog.
func BenchmarkDebugLog(b *testing.B) {
	log.SetOutput(io.Discard)
	b.ReportAllocs()
	counter := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		debuglog.Println("Hello, World!")
		counter++
	}
	// condition to avoid compiler optimization
	if counter != b.N {
		b.Fatalf("counter = %d, want %d", counter, b.N)
	}
}

// Benchmark loop without debuglog.
func BenchmarkNoDebugLog(b *testing.B) {
	b.ReportAllocs()
	counter := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter++
	}
	// condition to avoid compiler optimization
	if counter != b.N {
		b.Fatalf("counter = %d, want %d", counter, b.N)
	}
}
