BINARY := run

DIRS := data/client data/provider

DOMAIN ?= $(shell minikube ip).nip.io

build:
	go build -o $(BINARY) .

init:
	mkdir -p $(DIRS)

run: build init
	./run standalone eval

akash:
	(cd .. && make)

kube-init:
	( cd ../_run/multi && DOMAIN="$(DOMAIN)" make install-all )

kube-run: build
	./run kube eval -s host-base="$(DOMAIN)"

kube-init-run: kube-init kube-run

deps-install:
	go get github.com/buger/jsonparser
	go get github.com/ovrclk/gestalt/...

clean:
	rm -rf data $(BINARY)

reset: clean init build akash
rerun: reset run

.PHONY: build init run \
	akash deps-install \
	clean reset rerun
