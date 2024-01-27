package v2

import "testing"

func BenchmarkMemoryWaste(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MemoryWaste(100)
	}
}

func BenchmarkMemoryWasteParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			MemoryWaste(100)
		}
	})
}
