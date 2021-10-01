.PHONY: chart
.SILENT: ${.PHONY}

COMMIT := $(shell git rev-parse --short HEAD)
TAG := $(shell git describe --abbrev=0 | sed "s/v//")
BRANCH := $(shell git branch --show-current)

all: chart

clean:
ifeq ($(OS), Windows_NT)
	del /f /q gpts*.tgz
else
	rm -f gpts*.tgz
endif

chart: clean
ifeq ($(OS), Windows_NT)
	type gpts\Chart.template | sed "/^version/! s/\[\[GPTS_VERSION\]\]/${COMMIT}/" | sed "s/\[\[GPTS_VERSION\]\]/${TAG}-${BRANCH}/" >gpts\Chart.yaml
else
	cat gpts/Chart.template | sed "/^version/! s/\[\[GPTS_VERSION\]\]/${COMMIT}/" | sed "s/\[\[GPTS_VERSION\]\]/${TAG}-${BRANCH}/" >gpts/Chart.yaml
endif
	helm package gpts
ifeq ($(OS), Windows_NT)
	del /f /q gpts\Chart.yaml
else
	rm -f gpts/Chart.yaml
endif