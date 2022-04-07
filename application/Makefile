.PHONY: clean build build-static test coverage image
.SILENT: ${.PHONY}

GO_TEST := go test ./... -race -p 1
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_TAG := $(shell git describe --abbrev=0 | sed "s/v//")

BASE_FLAGS := -X 'main.version=${GIT_TAG}' -X 'main.gitCommit=${GIT_COMMIT}'

all: build

clean:
	rm -f *.out *.html gpts gpts.exe
	go clean -cache

run: clean
	go run  .

build: clean
	go build -ldflags "${BASE_FLAGS} -X 'main.binaryType=dynamic'" .

build-static: clean
	env CGO_ENABLED=0 go build -ldflags "${BASE_FLAGS} -X 'main.binaryType=static' -w -extldflags '-static'" .

test: clean
	${GO_TEST}

coverage: clean
	${GO_TEST} -v -covermode atomic -coverprofile cover.out
	go tool cover -html cover.out -o cover.html


# COMMANDS FOR LOCAL DEVELOPMENT ONLY

image: clean
	docker build . -t ghcr.io/icikowski/gpts:${GIT_COMMIT} --build-arg version=${GIT_TAG}-dev --build-arg gitCommit=${GIT_COMMIT}
	echo "Built Docker image for GPTS version \"${GIT_TAG}-dev\" with tag \"${GIT_COMMIT}\""
