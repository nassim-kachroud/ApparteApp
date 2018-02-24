.PHONY: install
install:
	go get -v ./

.PHONY: build
build:
	CGO_ENABLED=0 go build -o ./main -a -ldflags '-s' -installsuffix cgo main.go

.PHONY: test-unit
test-unit:
	go test -v `go list ./... | grep -v /vendor/` -tags=unit

.PHONY: test-integration
test-integration: 
	go test -v `go list ./... | grep -v /vendor/` -tags=integration