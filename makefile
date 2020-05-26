benchmark:
	go test -bench Profile -cpuprofile cpu.out click/multiclicker_test.go

pprof:
	go tool pprof cpu.out

clean:
	rm cpu.out click.test