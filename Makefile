VERSION=1.0.3

export CGO_ENABLED=0

build:
	go build -v -trimpath -ldflags "-s -w -X main.version=${VERSION}" -o ${GOPATH}/bin/alfred-workflow main.go
