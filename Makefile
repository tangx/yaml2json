BINARY:=yaml2json

GOOS:=$(shell go env GOOS)
GOARCH:=$(shell go env GOARCH)
VERSION=$(shell cat .version)



build.x:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(MAKE) build
	CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 $(MAKE) build
	CGO_ENABLED=0 GOOS=linux  GOARCH=arm64 $(MAKE) build

build:
	go build -o out/$(BINARY)-v$(VERSION)-$(GOOS)-$(GOARCH) .


tag:
	git tag v$(VERSION) && git push origin v$(VERSION)

clean:
	rm -rf out
