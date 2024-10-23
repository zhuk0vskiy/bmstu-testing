.PHONY: dep clean

build: dep
	@go build -o build/main

dep:
	@go mod download

clean:
	@rm -r ./build
