BIN := hcl2json
VERSION := $(shell git describe --tags --always --long --dirty)
LDFLAGS := -ldflags="-s -w -X main.version=$(VERSION)"
export CGO_ENABLED := 0

build:
	GOOS=linux   GOARCH=arm64 go build -o bin/$(BIN)_$(VERSION)_linux_arm64       $(LDFLAGS) main.go
	GOOS=linux   GOARCH=amd64 go build -o bin/$(BIN)_$(VERSION)_linux_amd64       $(LDFLAGS) main.go
	GOOS=darwin  GOARCH=amd64 go build -o bin/$(BIN)_$(VERSION)_darwin_amd64      $(LDFLAGS) main.go
	GOOS=windows GOARCH=amd64 go build -o bin/$(BIN)_$(VERSION)_windows_amd64.exe $(LDFLAGS) main.go
