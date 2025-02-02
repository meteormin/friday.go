PRJ_NAME=Friday.go
AUTHOR="Meteormin \(miniyu97@gmail.com\)"
PRJ_BASE=$(shell pwd)
PRJ_DESC=$(PRJ_NAME) Deployment and Development Makefile.\n Author: $(AUTHOR)

SUPPORTED_OS=linux darwin
SUPPORTED_ARCH=amd64 arm64

.DEFAULT: help
.SILENT:;

##help: helps (default)
.PHONY: help
help: Makefile
	echo ""
	echo " $(PRJ_DESC)"
	echo ""
	echo " Usage:"
	echo ""
	echo "	make {command}"
	echo ""
	echo " Commands:"
	echo ""
	sed -n 's/^##/	/p' $< | column -t -s ':' |  sed -e 's/^/ /'
	echo ""

##run mod={entrypoint: server, client, standalone}: run application
.PHONY: run
run:
	go run ./$(mod)/main.go

ldflags=-ldflags "-linkmode external -extldflags -static"

##build os={os: linux, darwin} arch={arch: amd64, arm64} mod={entrypoint: server, client, standalone}: build application
.PHONY: build
build:
ifeq ($(os), linux)
	CC=$(cc) CGO_ENABLED=1 GOOS=$(os) GOARCH=$(arch) go build $(ldflags) -x -o build/$(mod)-$(os)-$(arch) ./cmd/$(mod)/main.go
else
	CC=$(cc) CGO_ENABLED=1 GOOS=$(os) GOARCH=$(arch) go build -x -o build/$(mod)-$(os)-$(arch) ./cmd/$(mod)/main.go
endif

##release mod={entrypoint: server, client, standalone} tag={tag: v1.0.0}: release application
.PHONY: release
release:
	mkdir -p release/$(tag)
	$(foreach os, $(SUPPORTED_OS), \
		$(foreach arch, $(SUPPORTED_ARCH), \
			$(MAKE) build os=$(os) arch=$(arch) mod=$(mod)))
	cp build/$(mod)-$(os)-$(arch) release/$(tag)