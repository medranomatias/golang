remove:
	rm main *.pprof

build: remove
	go build main.go

run:
	./main ../bible.txt

view:
	go tool pprof cpu.pprof

mem:
	go tool pprof -http localhost:8080 mem.prof