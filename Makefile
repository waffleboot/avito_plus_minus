
.PHONY: bench run

mem: *.go
	go test -bench=. -benchmem -memprofile mem.out
	go tool pprof mem.out

cpu: *.go
	go test -bench=. -cpuprofile cpu.out
	go tool pprof cpu.out

run: *.go
	go run .
