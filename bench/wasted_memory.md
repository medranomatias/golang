```bash
go test ./... -bench=BenchmarkMemoryWaste
goos: darwin
goarch: amd64
pkg: github.com/medranomatias/benchmark
cpu: VirtualApple @ 2.50GHz
BenchmarkMemoryWaste-8           3125319               379.8 ns/op
PASS
ok      github.com/medranomatias/benchmark      3.589s
```

```bash
no$ go test ./... -bench=BenchmarkMemoryWaste -benchmem
goos: darwin
goarch: amd64
pkg: github.com/medranomatias/benchmark
cpu: VirtualApple @ 2.50GHz
BenchmarkMemoryWaste-8           3182074               371.5 ns/op          2040 B/op          8 allocs/op
PASS
ok      github.com/medranomatias/benchmark      1.944s

```

```bash
$ go test ./... -bench=BenchmarkMemoryWaste -benchmem -benchtime=1s -count=10
goos: darwin
goarch: amd64
pkg: github.com/medranomatias/benchmark
cpu: VirtualApple @ 2.50GHz
BenchmarkMemoryWaste-8           3165766               377.1 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           3200185               371.8 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           3429706               374.2 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           3207104               370.0 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           3180966               374.9 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           3051752               406.1 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           3103418               380.8 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           3107530               386.8 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           3161984               379.9 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           3141788               384.6 ns/op          2040 B/op          8 allocs/op
PASS
ok      github.com/medranomatias/benchmark      16.655s
```