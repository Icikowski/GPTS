FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gpts .

FROM alpine:latest  
LABEL maintainer "Piotr Icikowski <piotrekicikowski@gmail.com>"
LABEL org.opencontainers.image.title "GPTS - General Purpose Test Service"
LABEL org.opencontainers.image.description "Simple & easy to configure test service for serving content through HTTP"
LABEL org.opencontainers.image.authors "Piotr Icikowski"
LABEL org.opencontainers.image.vendor "Piotr Icikowski"
LABEL org.opencontainers.image.url "https://github.com/Icikowski/GPTS"
LABEL org.opencontainers.image.source "https://github.com/Icikowski/GPTS"
LABEL org.opencontainers.image.licenses "GPL-3.0-or-later"
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/gpts .
CMD ["./gpts"] 