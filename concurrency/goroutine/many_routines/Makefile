.DEFAULT_GOAL:=build

clean:
	rm -frv main *.pdf *.pprof *.out
build: clean
	go build -race main.go
pprof:
	go tool pprof main goroutine.pprof
trace:
	go tool trace -pprof=net trace.out > net.pprof
	go tool trace -pprof=sync trace.out > sync.pprof
	go tool trace -pprof=syscall trace.out > syscall.pprof
	go tool trace -pprof=sched trace.out > sched.pprof
	go tool trace trace.out