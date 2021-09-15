.PHONY: clean build test coverage image chart
.SILENT: ${.PHONY}

GO_TEST := go test ./... -race -p 1

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
	go build .

test: clean
	${GO_TEST}

coverage: clean
	${GO_TEST} -v -covermode atomic -coverprofile cover.out
	go tool cover -html cover.out -o cover.html


# COMMANDS FOR LOCAL DEVELOPMENT ONLY

image: clean
	docker build . -t ghcr.io/icikowski/gpts:$(shell git rev-parse --short HEAD)

chart: clean
ifeq ($(OS), Windows_NT)
	type ..\chart\gpts\Chart.template | sed "/^version/! s/\[\[GPTS_VERSION\]\]/$(shell git rev-parse --short HEAD)/" | sed "s/\[\[GPTS_VERSION\]\]/$(shell git describe --abbrev=0 | sed "s/v//")-dev/" >..\chart\gpts\Chart.yaml
else
	cat ../chart/gpts/Chart.template | sed "/^version/! s/\[\[GPTS_VERSION\]\]/$(shell git rev-parse --short HEAD)/" | sed "s/\[\[GPTS_VERSION\]\]/$(shell git describe --abbrev=0 | sed "s/v//")-dev/" >../chart/gpts/Chart.yaml
endif
	(cd ../chart && helm package gpts)
ifeq ($(OS), Windows_NT)
	del /f /q ..\chart\gpts\Chart.yaml
else
	rm -f ../chart/gpts/Chart.yaml
endif