# Build the manager binary
FROM golang:1.15 as builder

# copy root of the repo
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o cfgwriter main.go

FROM alpine:3.12.1
WORKDIR /
COPY --from=builder /workspace/cfgwriter .

ENTRYPOINT ["/cfgwriter"]
