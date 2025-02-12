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

SWAG_OUTPUT=./docs

##swag mod={entrypoint [friday]}: generate swagger docs
##: - output directory is `./docs`
.PHONY: swag
swag:
	swag init --parseDependency --parseInternal -g cmd/$(mod)/main.go --output api

WORKER_IMAGE=golang:1.22-alpine
OAS3_GENERATOR_DOCKER_IMAGE=openapitools/openapi-generator-cli:latest-release

##openapi3 mod={entrypoint [friday]}: Generate OAS3 from swaggo/swag output since that project doesn't support it
##: TODO. Remove this if V3 spec is ever returned from that project
.PHONY: openapi3
openapi3: swag
	@echo "[OAS3] Converting Swagger 2-to-3 (yaml)"
	@docker run --rm -v $(PWD)/api:/work $(OAS3_GENERATOR_DOCKER_IMAGE) \
	  generate -i /work/swagger.yaml -o /work/v3 -g openapi-yaml --minimal-update
	@docker run --rm -v $(PWD)/api/v3:/work $(WORKER_IMAGE) \
	  sh -c "rm -rf /work/.openapi-generator"
	@echo "[OAS3] Copying openapi-generator-ignore (json)"
	@docker run --rm -v $(PWD)/api/v3:/work $(WORKER_IMAGE) \
	  sh -c "cp -f /work/.openapi-generator-ignore /work/openapi"
	@echo "[OAS3] Converting Swagger 2-to-3 (json)"
	@docker run --rm -v $(PWD)/api:/work $(OAS3_GENERATOR_DOCKER_IMAGE) \
	  generate -s -i /work/swagger.json -o /work/v3/openapi -g openapi --minimal-update
	@echo "[OAS3] Cleaning up generated files"
	@docker run --rm -v $(PWD)/api/v3:/work $(WORKER_IMAGE) \
	  sh -c "mv -f /work/openapi/openapi.json /work ; mv -f /work/openapi/openapi.yaml /work ; rm -rf /work/openapi"



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
	CC=$(cc) CGO_ENABLED=1 GOOS=$(os) GOARCH=$(arch) go build $(LDFLAGS) -o build/$(mod)-$(os)-$(arch) ./cmd/$(mod)/main.go
else
	CC=$(cc) CGO_ENABLED=1 GOOS=$(os) GOARCH=$(arch) go build -o build/$(mod)-$(os)-$(arch) ./cmd/$(mod)/main.go
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