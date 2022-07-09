.PHONY: clean build
.SILENT: ${.PHONY}

COMMIT := $(shell git rev-parse --short HEAD)
TAG := $(shell git describe --abbrev=0 | sed "s/v//")

all: chart

clean:
	rm -f gpts*.tgz

build: clean
	helm package --app-version ${COMMIT} --version ${TAG}-dev gpts
	echo "Built Helm chart \"${TAG}-dev\" for GPTS image with tag \"${COMMIT}\""

lint: clean
	helm lint gpts
