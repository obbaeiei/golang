package golang

import "testing"

// go test -run BenchmarkMaxAPU -bench . -benchtime 3s -benchmem -cpuprofile cpu.out
// go tool pprof try.test cpu.out
func BenchmarkMaxAPU(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		loop(10000)
	}
}
