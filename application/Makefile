.PHONY: clean build test coverage image chart
.SILENT: ${.PHONY}

GO_TEST := go test ./... -race -p 1
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_TAG := $(shell git describe --abbrev=0 | sed "s/v//")

BASE_FLAGS := -X 'main.version=${GIT_TAG}' -X 'main.gitCommit=${GIT_COMMIT}'

ifeq ($(OS), Windows_NT)
	BUILD_BEG := set CGO_ENABLED=0 &&
else
	BUILD_BEG := env CGO_ENABLED=0
endif

all: build

clean:
ifeq ($(OS), Windows_NT)
	del /f /q *.out *.html gpts gpts.exe
else
	rm -f *.out *.html gpts gpts.exe
endif
	go clean -cache

run: clean
	go run  .

build: clean
	go build -ldflags "${BASE_FLAGS} -X 'main.binaryType=dynamic'" .

build-static: clean
	${BUILD_BEG} go build -ldflags "${BASE_FLAGS} -X 'main.binaryType=static' -w -extldflags '-static'" .

test: clean
	${GO_TEST}

coverage: clean
	${GO_TEST} -v -covermode atomic -coverprofile cover.out
	go tool cover -html cover.out -o cover.html


# COMMANDS FOR LOCAL DEVELOPMENT ONLY

image: clean
	docker build . -t ghcr.io/icikowski/gpts:${GIT_COMMIT} --build-arg version=${GIT_TAG}-dev --build-arg gitCommit=${GIT_COMMIT}
