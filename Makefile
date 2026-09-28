GOPATH ?= /usr/local/go/bin/go

.PHONY: test coverage html clean

test:
	$(GOPATH) test -v ./...

coverage:
	$(GOPATH) test -v -coverprofile=coverage.txt -covermode=atomic ./...
	$(GOPATH) tool cover -func=coverage.txt

html: coverage
	$(GOPATH) tool cover -html=coverage.txt -o coverage.html
	@echo "Coverage HTML generated at coverage.html"

clean:
	rm -f coverage.txt coverage.html
