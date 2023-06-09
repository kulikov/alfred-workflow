SUFFIX?=""
TARGET_OS=linux darwin
TARGET_ARCH=amd64

export CGO_ENABLED=0

build:
	go build -v -trimpath -ldflags "-s -w -X main.version=${VERSION}" -o bin/alfred-workflow${SUFFIX} main.go
