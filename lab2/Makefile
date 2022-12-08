GO := $(shell find . -name "*.go")

bin/server: $(GO)
	go build -o bin/server prog.go

.PHONY: clean run

run_server: bin
	./bin/server

clean:
	rm -rf bin pb
