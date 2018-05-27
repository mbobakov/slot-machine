.PHONY: run linux test lint clean integration
build: test
	go build -o slot-machine *.go
run:
	go run *.go
linux:
	GOOS=linux  go build -o slot-machine-linux *.go
test:
	go test -v ./...
lint:
	gometalinter --vendor --enable-all --deadline=2m  --line-length=150 ./...
clean:
	rm -f slot-machine*