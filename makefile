vpath %.go ./click ./parallel
gofiles = *.go
.PHONY : clean

build: $(gofiles) parallel/$(gofiles) click/$(gofiles)
	go build

benchmark:
	go test -bench Profile -cpuprofile cpu.out click/multiclicker_test.go

pprof:
	go tool pprof cpu.out

clean:
	-rm cpu.out click.test demo