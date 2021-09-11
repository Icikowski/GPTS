.PHONY: clean build test coverage image
.SILENT: ${.PHONY}

GO_TEST := go test ./... -race -p 1

all: build

clean:
ifeq ($(OS), Windows_NT)
	del /f /s /q *.out *.html gpts gpts.exe
else
	rm -f *.out *.html gpts gpts.exe
endif
	go clean -cache

run: clean
	go run  .

build: clean
	go build .

test: clean
	${GO_TEST}

coverage: clean
	${GO_TEST} -v -covermode atomic -coverprofile cover.out
	go tool cover -html cover.out -o cover.html

image: clean
	docker build . -t ghcr.io/icikowski/gpts:$(shell git rev-parse --short HEAD)
