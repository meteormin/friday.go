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

##run mod={entrypoint [friday]}: run application
.PHONY: run
run:
	go run ./cmd/$(mod)/main.go


# OS와 ARCH가 정의되어 있지 않으면 기본값을 설정합니다.
# uname -s는 OS 이름(예: Linux, Darwin 등)을 반환하고, tr를 통해 소문자로 변환합니다.
OS ?= $(shell uname -s | tr '[:upper:]' '[:lower:]')
# uname -m은 아키텍처 정보를 반환합니다. (예: x86_64, arm64 등)
ARCH ?= $(shell uname -m)
LDFLAGS=-ldflags "-linkmode external -extldflags -static"

##build os={os [linux, darwin]} arch={arch [amd64, arm64]} mod={entrypoint [friday]}: build application
.PHONY: build
build: os ?= $(OS)
build: arch ?= $(ARCH)
build:
	echo "Building $(mod) for $(os)-$(arch)"
ifeq ($(os), linux)
	CC=$(cc) CGO_ENABLED=1 GOOS=$(os) GOARCH=$(arch) go build $(LDFLAGS) -x -o build/$(mod)-$(os)-$(arch) ./cmd/$(mod)/main.go
else
	CC=$(cc) CGO_ENABLED=1 GOOS=$(os) GOARCH=$(arch) go build -x -o build/$(mod)-$(os)-$(arch) ./cmd/$(mod)/main.go
endif

##release mod={entrypoint [friday]} tag={tag [v1.0.0]}: release application
.PHONY: release
release:
	mkdir -p release/$(tag)
	$(foreach os, $(SUPPORTED_OS), \
		$(foreach arch, $(SUPPORTED_ARCH), \
			$(MAKE) build os=$(os) arch=$(arch) mod=$(mod)))
	cp build/$(mod)-$(os)-$(arch) release/$(tag)
	cp build/config.yml release/config.yml

##clean: clean application
.PHONY: clean
clean:
	rm -rf build/*