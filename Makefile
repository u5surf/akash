PROTO_FILES  = $(wildcard types/*.proto)
PROTOC_FILES = $(patsubst %.proto,%.pb.go, $(PROTO_FILES))

BINS       := akash akashd
IMAGE_BINS := _build/akash _build/akashd

GO := GO111MODULE=on go

IMAGE_BUILD_ENV = GOOS=linux GOARCH=amd64

BUILD_FLAGS = -ldflags \
							"-X github.com/ovrclk/akash/version.version=$(shell git rev-parse --abbrev-ref HEAD) \
					     -X github.com/ovrclk/akash/version.commit=$(shell git rev-parse HEAD) \
					     -X github.com/ovrclk/akash/version.date=$(shell date +%Y-%m-%dT%H:%M:%S%z)"

all: build bins

bins: $(BINS)

build:
	$(GO) build ./...

akash:
	$(GO) build $(BUILD_FLAGS) ./cmd/akash

akashd:
	$(GO) build $(BUILD_FLAGS) ./cmd/akashd

image-bins:
	$(IMAGE_BUILD_ENV) $(GO) build $(BUILD_FLAGS) -o _build/akash  ./cmd/akash
	$(IMAGE_BUILD_ENV) $(GO) build $(BUILD_FLAGS) -o _build/akashd ./cmd/akashd

image: image-bins
	docker build --rm            \
		-t ovrclk/akash:latest     \
		-f _build/Dockerfile.akash \
		_build
	docker build --rm             \
		-t ovrclk/akashd:latest     \
		-f _build/Dockerfile.akashd \
		_build

install:
	$(GO) install $(BUILD_FLAGS) ./cmd/akash
	$(GO) install $(BUILD_FLAGS) ./cmd/akashd

release:
	docker run --rm --privileged \
	-v $(PWD):/go/src/github.com/ovrclk/akash \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-w /go/src/github.com/ovrclk/akash \
	-e GITHUB_TOKEN \
	-e DOCKER_USERNAME \
	-e DOCKER_PASSWORD \
	-e DOCKER_REGISTRY \
	goreleaser/goreleaser release

image-minikube:
	eval $$(minikube docker-env) && make image

test:
	$(GO) test ./...

test-nocache:
	$(GO) test -count=1 ./...

test-full:
	$(GO) test -race ./...

test-lint:
	golangci-lint run

lintdeps-install:
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint

test-vet:
	$(GO) vet ./...

deps-install:
	$(GO) mod download

devdeps-install:
	$(GO) install github.com/gogo/protobuf/protoc-gen-gogo
	$(GO) install github.com/vektra/mockery/.../
	$(GO) install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	$(GO) install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

test-integration: $(BINS)
	(cd _integration && make clean run)

integrationdeps-install:
	(cd _integration && make deps-install)

gentypes: $(PROTOC_FILES)

kubetypes:
	vendor/k8s.io/code-generator/generate-groups.sh all \
  	github.com/ovrclk/akash/pkg/client github.com/ovrclk/akash/pkg/apis \
  	akash.network:v1

%.pb.go: %.proto
	protoc -I. \
		-Ivendor -Ivendor/github.com/gogo/protobuf/protobuf \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogo_out=plugins=grpc:. $<
	protoc -I. \
		-Ivendor -Ivendor/github.com/gogo/protobuf/protobuf \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. $<
	protoc -I. \
		-Ivendor -Ivendor/github.com/gogo/protobuf/protobuf \
		-Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. $<

mocks:
	mockery -case=underscore -dir query                 -output query/mocks                 -name Client
	mockery -case=underscore -dir txutil                -output txutil/mocks                -name Client
	mockery -case=underscore -dir app/market            -output app/market/mocks            -name Client
	mockery -case=underscore -dir app/market            -output app/market/mocks            -name Engine
	mockery -case=underscore -dir app/market            -output app/market/mocks            -name Facilitator
	mockery -case=underscore -dir marketplace           -output marketplace/mocks           -name Handler
	mockery -case=underscore -dir provider              -output provider/mocks              -name StatusClient
	mockery -case=underscore -dir provider/cluster      -output provider/cluster/mocks      -name Client
	mockery -case=underscore -dir provider/cluster      -output provider/cluster/mocks      -name Cluster
	mockery -case=underscore -dir provider/cluster      -output provider/cluster/mocks      -name Deployment
	mockery -case=underscore -dir provider/cluster      -output provider/cluster/mocks      -name Reservation
	mockery -case=underscore -dir provider/cluster/kube -output provider/cluster/kube/mocks -name Client
	mockery -case=underscore -dir provider/manifest     -output provider/manifest/mocks     -name Handler


gofmt:
	find . -not -path './vendor*' -name '*.go' -type f | \
		xargs gofmt -s -w

docs:
	(cd _docs/dot && make)

clean:
	rm -f $(BINS) $(IMAGE_BINS)

.PHONY: all bins build \
	akash akashd \
	image image-bins \
	test test-nocache test-full \
	deps-install devdeps-install \
	test-integraion integrationdeps-install \
	test-lint lintdeps-install \
	test-vet \
	mocks \
	gofmt \
	docs \
	clean \
	kubetypes gentypes $(PROTO_FILES) \
	install
