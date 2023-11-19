.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -e -v

.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...

.PHONY: test
test:
	go test -v -race -buildvcs ./...

.PHONY: build
build:
	wails build -platform windows/amd64

# Wails CLI handles building for the 'run' command
.PHONY: run
run:
	wails dev
