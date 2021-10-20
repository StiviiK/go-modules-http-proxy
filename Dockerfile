# Builder
FROM golang:1.17 as builder
WORKDIR /app

# Add module files
ADD go.mod .
ADD go.sum .

# Fetch dependencies.
RUN go mod download
RUN go mod verify

# Add the source code
ADD . .

# Build the binary.
RUN GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags="-X main.version=${VERSION} -X main.Version=$(go version | cut -d " " -f 3) -X main.BuildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
    -a -o \
    /usr/local/bin/go-modules-http-proxy

# Runner
FROM gcr.io/distroless/base

# Copy our static executable.
COPY --from=builder /usr/local/bin/go-modules-http-proxy /usr/local/bin/modulesproxy

CMD ["/usr/local/bin/modulesproxy", "serve"]