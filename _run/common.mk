AKASH_ROOT := ../..

DATA_ROOT = data
NODE_ROOT = $(DATA_ROOT)/node

all:
	(cd $(AKASH_ROOT) && make bins)
build:
	(cd $(AKASH_ROOT) && make build)
akash:
	(cd $(AKASH_ROOT) && make akash)
akashd:
	(cd $(AKASH_ROOT) && make akashd)
image-minikube:
	(cd $(AKASH_ROOT) && make image-minikube)

clean:
	rm -rf $(DATA_ROOT)
