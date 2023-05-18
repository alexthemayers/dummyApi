.phony: run build

run: build
	./dummyApi
build:
	make clean && make dummyApi

dummyApi:
	go mod tidy
	goimports -w .
	go vet ./...
	go build

clean:
	rm -rf ./dummyApi
