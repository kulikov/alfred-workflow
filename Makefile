VERSION=1.0.2
TARGET_OS=darwin
TARGET_ARCH=arm64

export CGO_ENABLED=0

build:
	for GOOS in ${TARGET_OS}; do \
		for GOARCH in ${TARGET_ARCH}; do \
			GOOS=$$GOOS GOARCH=$$GOARCH go build -v -trimpath -ldflags "-s -w -X main.version=${VERSION}" -o bin/alfred-workflow-v${VERSION}-$$GOOS-$$GOARCH main.go; \
		done \
	done \
