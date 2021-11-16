FROM golang:latest AS builder
ARG version
ARG gitCommit
ENV BUILD_VERSION=${version:-unknown}
ENV BUILD_GIT_COMMIT=${gitCommit:-unknown}
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download -x
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gpts -ldflags "-X 'main.version=${BUILD_VERSION}' -X 'main.gitCommit=${BUILD_GIT_COMMIT}' -X 'main.binaryType=static' -w -extldflags '-static'" .

FROM gcr.io/distroless/static:nonroot
LABEL maintainer "Piotr Icikowski <piotrekicikowski@gmail.com>"
LABEL org.opencontainers.image.title "GPTS - General Purpose Test Service"
LABEL org.opencontainers.image.description "Simple & easy to configure test service for serving content through HTTP"
LABEL org.opencontainers.image.authors "Piotr Icikowski"
LABEL org.opencontainers.image.vendor "Piotr Icikowski"
LABEL org.opencontainers.image.url "https://github.com/Icikowski/GPTS"
LABEL org.opencontainers.image.source "https://github.com/Icikowski/GPTS"
LABEL org.opencontainers.image.licenses "GPL-3.0-or-later"
WORKDIR /
COPY --from=builder /app/gpts .
USER 65532:65532
ENTRYPOINT ["/gpts"]