.PHONY: chart
.SILENT: ${.PHONY}

COMMIT := $(shell git rev-parse --short HEAD)
TAG := $(shell git describe --abbrev=0 | sed "s/v//")
BRANCH := $(shell git branch --show-current)

all: chart

clean:
	rm -f gpts*.tgz gpts/Chart.yaml

build: clean
	cp gpts/Chart.template gpts/Chart.yaml
	sed -i "/^version/! s/\[\[GPTS_VERSION\]\]/${COMMIT}/" gpts/Chart.yaml 
	sed -i "s/\[\[GPTS_VERSION\]\]/${TAG}-dev/" gpts/Chart.yaml
	helm package gpts
	echo "Built Helm chart \"${TAG}-dev\" for GPTS image with tag \"${COMMIT}\""