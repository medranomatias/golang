clean:
	rm -fr main *.pprof *.pdf
build: clean
	go build main.go
run:
	./main ../bible.txt
time: build
	time ./main ../bible.txt
wc:
	time wc ../bible.txt
pprof:
	go tool pprof cpu.pprof
web:
	go tool pprof -http localhost:8080 main.go cpu.pprof