PATH := ${CURDIR}/bin:$(PATH)

go_exe = $(shell go env GOEXE)

vendor:
	dep ensure

bin/%:
	go build -o $@ -tags netgo -ldflags "-w" ./cmd/$(basename $*)

.PHONY: clean
clean:
	rm -rf bin
