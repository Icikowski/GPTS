.PHONY: clean build test coverage image
.SILENT: ${.PHONY}

GO_TEST := go test ./... -race -p 1

all: build

clean:
	rm -f *.out *.html gpts gpts.exe
	go clean -cache

build: clean
	go build .

test: clean
	${GO_TEST}

coverage: clean
	${GO_TEST} -v -covermode atomic -coverprofile cover.out
	go tool cover -html cover.out -o cover.html

image: clean
	docker build . -t ghcr.io/icikowski/gpts:$(shell cat Dockerfile | grep '# Version' | awk '{{ print $$3 }}')
