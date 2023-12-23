.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -e -v

.PHONY: test
test:
	go test -v -race -buildvcs ./...

# builds based on current platform as set by GOOS/GOARCH etc
.PHONY: build
build:
	wails build -debug

.PHONY: build-windows
build-windows:
	go generate
	GOOS=windows wails build -platform windows/amd64 -clean -debug

# Wails CLI handles building for the 'run' command
.PHONY: run
run:
	./scripts/dev.sh
