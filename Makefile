PATH := ${CURDIR}/bin:$(PATH)

go_exe = $(shell go env GOEXE)

vendor:
	dep ensure

bin/mockery$(go_exe): vendor
	go build -o $@ ./vendor/github.com/vektra/mockery/cmd/mockery

bin/%:
	go build -o $@ -tags netgo -ldflags "-w" ./cmd/$(basename $*)

.PHONY: generate
generate: bin/mockery$(go_exe)
	go generate ./...

.PHONY: clean
clean:
	rm -rf bin
