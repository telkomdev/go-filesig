.PHONY : test format build clean

ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

test:
	$(foreach pkg, $(ALL_PACKAGES),\
	go test -race -v $(pkg);)

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

build:
	@echo "building binary"
	@go build -o app  github.com/telkomdev/go-filesig/cmd/app

build-webapp:
	@echo "building binary"
	@go build -o webapp  github.com/telkomdev/go-filesig/cmd/webapp

clean:
	@echo "cleaning unused file"
	rm -rf app \
	rm -rf webapp \
	&& rm -rf *.txt